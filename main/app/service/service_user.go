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
	userData, err := ctx.repo.UserLoginRepo(username)
	if err != nil {
		// 记录错误日志
		global.GVA_LOGGER.Error("用户登录查询失败", zap.String("err", err.Error()))
		return nil, er.JZError("登录失败")
	}
	isOk := byt.ComparePassword(userData.Password, password)
	if !isOk {
		return nil, er.JZError("密码错误")
	}
	// 查询头像
	avatarData, err := ctx.repo.GetUserAvatarRepo(userData.UID)
	if err != nil {
		global.GVA_LOGGER.Error("用户头像查询失败", zap.String("err", err.Error()))
		return nil, er.JZError("登录失败")
	}
	// 生成token
	jwtToken, err := auth.GenerateToken(avatarData.IID, username, userData.UID)
	if err != nil {
		global.GVA_LOGGER.Error("生成token失败", zap.String("err", err.Error()))
		return nil, er.JZError("登录失败")
	}
	// 加密token
	stringToken, err := ead.EncryptWithPublicKeyOAEP(global.GVA_PUBLIC_KEY, jwtToken)
	if err != nil {
		global.GVA_LOGGER.Error("加密token失败", zap.String("err", err.Error()))
		return nil, er.JZError("登录失败")
	}

	// 拼接头像链接
	conf := global.GVA_CONFIG.Minio
	var avatarBuilder strings.Builder
	avatarBuilder.WriteString(conf.FilePath)
	avatarBuilder.WriteString("/")
	avatarBuilder.WriteString(avatarData.Path)
	avatarUrl := avatarBuilder.String()

	return map[string]interface{}{
		"uid":      userData.UID,
		"image":    avatarUrl,
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
