package repo

import (
	"gorm.io/gorm"
	"jingzhe-bg/main/app/model"
	"jingzhe-bg/main/global"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{db: global.GVA_DB}
}

// 登录查询
func (ctx *UserRepo) UserLoginRepo(username string) (*model.UserModel, error) {
	var user model.UserModel
	if select_err := ctx.db.
		Select("uid, password, nickname").
		Where("username = ? AND status = 1", username).
		Find(&user).Error; select_err != nil {
		return nil, select_err
	}

	return &user, nil
}

// 查询头像链接
func (ctx *UserRepo) GetUserAvatarRepo(uid string) (*model.UserImageModel, error) {
	var avatar model.UserImageModel
	if select_err := ctx.db.
		Select("id, minio_path").
		Where("uid = ? AND is_deleted = 0 AND is_avatar = 1", uid).
		First(&avatar).Error; select_err != nil {
		return nil, select_err
	}
	return &avatar, nil
}
