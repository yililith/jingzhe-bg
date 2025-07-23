package repo

import (
	"gorm.io/gorm"
	"jingzhe-bg/main/app/dto"
	"jingzhe-bg/main/app/model"
	"jingzhe-bg/main/global"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{db: global.GVA_DB}
}

// UserLoginRepo
//
//	@Description: 查询用户uid,密码和昵称
//	@receiver ctx
//	@param username
//	@return *model.UserModel
//	@return er
func (ctx *UserRepo) UserLoginRepo(username string) (*model.UserModel, error) {
	var user model.UserModel
	if err := ctx.db.
		Select("uid, password, nickname").
		Where("username = ? AND status = 1", username).
		Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserAvatarRepo
//
//	@Description: 查询用户头像链接
//	@receiver ctx
//	@param uid
//	@return *model.UserImageModel
//	@return er
func (ctx *UserRepo) GetUserAvatarRepo(uid int64) (*model.UserImageModel, error) {
	var avatar model.UserImageModel
	if err := ctx.db.
		Select("id, minio_path").
		Where("uid = ? AND is_deleted = 0 AND is_avatar = 1", uid).
		First(&avatar).Error; err != nil {
		return nil, err
	}
	return &avatar, nil
}

// HasUserRepo
//
//	@Description: 查询用户是否存在
//	@receiver ctx
//	@param username
//	@return int64
//	@return er
func (ctx *UserRepo) HasUserRepo(username string) (int64, error) {
	var count int64
	tx := ctx.db.Model(&model.UserModel{}).
		Where("username = ?", username).
		Count(&count)

	if tx.Error != nil {
		return 1, tx.Error
	}
	return count, nil
}

// CreateNewUserRepo
//
//	@Description: 创建用户
//	@receiver ctx
//	@param id
//	@return er
func (ctx *UserRepo) CreateNewUserRepo(newUser *dto.CreateUserDto) error {
	begin := ctx.db.Begin()

	if begin.Error != nil {
		begin.Rollback()
		return begin.Error
	}

	begin.Create(&model.UserModel{
		Username: newUser.Username,
		Password: newUser.Password,
		Nickname: newUser.Nickname,
		Grade:    2,
		Status:   1,
	})

	if err := begin.Commit().Error; err != nil {
		begin.Rollback()
		return err
	}

	return nil

}
