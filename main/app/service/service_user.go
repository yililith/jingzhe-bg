package service

import (
	"bytes"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"io"
	"jingzhe-bg/main/app/dto"
	"jingzhe-bg/main/app/repo"
	"jingzhe-bg/main/global"
	"jingzhe-bg/main/utils/auth"
	"jingzhe-bg/main/utils/byt"
	"jingzhe-bg/main/utils/ead"
	"jingzhe-bg/main/utils/er"
	"jingzhe-bg/main/utils/oss"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		repo: repo.NewUserRepo(),
	}
}

// UserLoginService
//
//	@Description: 登录
//	@receiver ctx
//	@param username
//	@param password
//	@return map[string]interface{}
//	@return er
func (ctx *UserService) UserLoginService(username string, password string) (map[string]interface{}, error) {
	// 密码查询验证
	userData, err := ctx.repo.UserLoginRepo(username)
	if err != nil {
		return nil, err
	}
	isOk := byt.ComparePassword(userData.Password, password)
	if !isOk {
		return nil, er.JZError("密码错误")
	}
	// 查询头像
	avatarData, err := ctx.repo.GetUserAvatarRepo(userData.UID)
	if err != nil {
		return nil, er.JZError("登录失败")
	}
	// 生成token
	jwtToken, err := auth.GenerateToken(avatarData.ETag, username, userData.UID)
	if err != nil {
		return nil, er.JZError("登录失败")
	}
	// 加密token
	stringToken, err := ead.EncryptWithPublicKeyOAEP(global.GVA_PUBLIC_KEY, jwtToken)
	if err != nil {
		return nil, er.JZError("登录失败")
	}

	// 获取头像链接加密链接
	url, err := oss.GeneratePresignedURL(
		global.GAV_MINIO_CLIENT,
		global.GVA_CONFIG.Minio.BucketName,
		avatarData.ObjectName,
		2,
	)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"uid":      userData.UID,
		"image":    url,
		"token":    "Bearer " + stringToken,
		"nickname": userData.Nickname,
	}, nil
}

// CreateUserService
//
//	@Description: 创建新用户service
//	@receiver ctx
//	@param newUser
//	@return er
func (ctx *UserService) CreateUserService(newUser *dto.CreateUserDto) error {
	// 字段校验
	if err := global.GVA_VALIDATOR.Struct(newUser); err != nil {
		return err
	}
	// 查看用户是否存在
	userNum, err := ctx.repo.HasUserRepo(newUser.Username)
	if err != nil {
		global.GVA_LOGGER.Error(
			"用户查询失败",
			zap.String("err", err.Error()))
		return er.JZError("创建用户失败")
	}
	if userNum > 0 {
		return er.JZError("账号已存在")
	}

	// 加密密码
	password, err := byt.HashPassword(newUser.Password)
	if err != nil {
		global.GVA_LOGGER.Error(
			"加密密码失败",
			zap.String("err", err.Error()),
		)
		return er.JZError("创建用户失败")
	}

	// 创建用户

	if err := ctx.repo.CreateNewUserRepo(&dto.CreateUserDto{
		Username: newUser.Username,
		Password: password,
		Nickname: newUser.Nickname,
	}); err != nil {
		global.GVA_LOGGER.Error("创建用户失败", zap.String("err", err.Error()))
		return er.JZError("创建用户失败")
	}
	return nil
}

// PutUserImageService
//
//	@Description: 上传图片
//	@receiver ctx
//	@param file
//	@param size
//	@return error
func (ctx *UserService) PutUserAvatarService(uid uint64, file multipart.File, header *multipart.FileHeader) error {
	const maxFileSize = 5 << 20 // 5MB
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	// 校验文件大小
	if header.Size > maxFileSize {
		return er.JZError("文件大小超出限制（最大 5MB）")
	}

	// 校验扩展名
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedExts[ext] {
		return er.JZError("文件格式不支持，仅支持 jpg/jpeg/png/webp")
	}

	// 读取文件到内存，避免流偏移问题
	data, err := io.ReadAll(file)
	if err != nil {
		global.GVA_LOGGER.Error("文件读取失败", zap.String("err", err.Error()))
		return er.JZError("文件读取失败")
	}

	// 生成文件名
	newFileName := time.Now().Format("2006/01/02/150405") + ".webp"

	var objectInfo *minio.UploadInfo

	// webp 直接上传
	if ext == ".webp" {
		objectInfo, err = oss.PutObject(
			global.GAV_MINIO_CLIENT,
			global.GVA_CONFIG.Minio.BucketName,
			newFileName,
			bytes.NewReader(data),
			int64(len(data)),
		)
		if err != nil {
			return err
		}
	} else {
		// 解码图片
		img, err := imaging.Decode(bytes.NewReader(data))
		if err != nil {
			global.GVA_LOGGER.Error("文件解码失败", zap.String("err", err.Error()))
			return er.JZError("文件解码失败")
		}

		// 编码为 webp
		var buf bytes.Buffer
		defer buf.Reset()

		if err := webp.Encode(&buf, img, &webp.Options{
			Lossless: true,
			Quality:  100,
		}); err != nil {
			global.GVA_LOGGER.Error("文件编码失败", zap.String("err", err.Error()))
			return er.JZError("文件编码失败")
		}

		objectInfo, err = oss.PutObject(
			global.GAV_MINIO_CLIENT,
			global.GVA_CONFIG.Minio.BucketName,
			newFileName,
			bytes.NewReader(buf.Bytes()),
			int64(buf.Len()),
		)
		if err != nil {
			return err
		}
	}

	if err = ctx.repo.PutUserImageRepo(uid, "图片", objectInfo, true); err != nil {
		return err
	}

	return nil
}
