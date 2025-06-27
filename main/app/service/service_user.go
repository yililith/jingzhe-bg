package service

import (
	"errors"
	"jingzhe-bg/main/app/repo"
	"jingzhe-bg/main/internal/rsa"
	"jingzhe-bg/main/utils/auth"
	"jingzhe-bg/main/utils/byt"
	"jingzhe-bg/main/utils/ead"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		repo: repo.NewUserRepo(),
	}
}

// 登录接口 service
func (ctx *UserService) UserLoginService(username string, password string) (map[string]interface{}, error) {
	// 密码查询验证
	user_data, user_repo_err := ctx.repo.UserLoginRepo(username)
	if user_repo_err != nil {
		return nil, user_repo_err
	}
	isOk := byt.ComparePassword(user_data.Password, password)
	if !isOk {
		return nil, errors.New("密码错误")
	}
	// 查询头像
	avatar_data, avatat_repo_err := ctx.repo.GetUserAvatarRepo(user_data.UID)
	if avatat_repo_err != nil {
		return nil, avatat_repo_err
	}
	// 生成token
	jwtToken, token_err := auth.GenerateToken(avatar_data.ID, user_data.UID, username)
	if token_err != nil {
		return nil, token_err
	}
	// 加密token
	stringToken, en_err := ead.EncryptWithPublicKeyOAEP(rsa.PublicKey, jwtToken)
	if en_err != nil {
		return nil, en_err
	}
	return map[string]interface{}{
		"uid":      user_data.UID,
		"image":    avatar_data.MinioPath,
		"token":    "Bearer " + stringToken,
		"nickname": user_data.Nickname,
	}, nil
}
