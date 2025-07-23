package service

import (
	"go.uber.org/zap"
	"jingzhe-bg/main/app/dto"
	"jingzhe-bg/main/app/repo"
	"jingzhe-bg/main/global"
	"jingzhe-bg/main/utils/auth"
	"jingzhe-bg/main/utils/byt"
	"jingzhe-bg/main/utils/ead"
	"jingzhe-bg/main/utils/er"
	"strings"
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
	user_data, user_repo_err := ctx.repo.UserLoginRepo(username)
	if user_repo_err != nil {
		// 记录错误日志
		global.GVA_LOGGER.Error("用户登录查询失败", zap.String("err", user_repo_err.Error()))
		return nil, er.JZError("登录失败")
	}
	isOk := byt.ComparePassword(user_data.Password, password)
	if !isOk {
		return nil, er.JZError("密码错误")
	}
	// 查询头像
	avatar_data, avatar_repo_err := ctx.repo.GetUserAvatarRepo(user_data.UID)
	if avatar_repo_err != nil {
		global.GVA_LOGGER.Error("用户头像查询失败", zap.String("err", avatar_repo_err.Error()))
		return nil, er.JZError("登录失败")
	}
	// 生成token
	jwtToken, token_err := auth.GenerateToken(avatar_data.IID, username, user_data.UID)
	if token_err != nil {
		global.GVA_LOGGER.Error("生成token失败", zap.String("err", token_err.Error()))
		return nil, er.JZError("登录失败")
	}
	// 加密token
	stringToken, en_err := ead.EncryptWithPublicKeyOAEP(global.GVA_PUBLIC_KEY, jwtToken)
	if en_err != nil {
		global.GVA_LOGGER.Error("加密token失败", zap.String("err", en_err.Error()))
		return nil, er.JZError("登录失败")
	}

	// 拼接头像链接
	conf := global.GVA_CONFIG.Minio
	var avatarBuilder strings.Builder
	avatarBuilder.WriteString(conf.FilePath)
	avatarBuilder.WriteString("/")
	avatarBuilder.WriteString(avatar_data.Path)
	avatarUrl := avatarBuilder.String()

	return map[string]interface{}{
		"uid":      user_data.UID,
		"image":    avatarUrl,
		"token":    "Bearer " + stringToken,
		"nickname": user_data.Nickname,
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
	if vaodator_err := global.GVA_VALIDATOR.Struct(newUser); vaodator_err != nil {
		return vaodator_err
	}
	// 查看用户是否存在
	userNum, count_err := ctx.repo.HasUserRepo(newUser.Username)
	if count_err != nil {
		global.GVA_LOGGER.Error(
			"用户查询失败",
			zap.String("err", count_err.Error()))
		return er.JZError("创建用户失败")
	}
	if userNum > 0 {
		return er.JZError("账号已存在")
	}

	// 加密密码
	password, hash_error := byt.HashPassword(newUser.Password)
	if hash_error != nil {
		global.GVA_LOGGER.Error(
			"加密密码失败",
			zap.String("err", hash_error.Error()),
		)
		return er.JZError("创建用户失败")
	}

	// 创建用户

	create_err := ctx.repo.CreateNewUserRepo(&dto.CreateUserDto{
		Username: newUser.Username,
		Password: password,
		Nickname: newUser.Nickname,
	})
	if create_err != nil {
		global.GVA_LOGGER.Error("创建用户失败", zap.String("err", create_err.Error()))
		return er.JZError("创建用户失败")
	}
	return nil
}
