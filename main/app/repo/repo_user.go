package repo

import (
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"jingzhe-bg/main/app/dto"
	"jingzhe-bg/main/app/model"
	"jingzhe-bg/main/global"
	"jingzhe-bg/main/utils/er"
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
		Where(&model.UserModel{
			Username: username,
			Status:   1,
		}).Find(&user).Error; err != nil {
		global.GVA_LOGGER.Error("登录查询失败", zap.String("err", err.Error()))
		return nil, er.JZError("查询失败")
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
func (ctx *UserRepo) GetUserAvatarRepo(uid uint64) (*model.UserImageModel, error) {
	var avatar model.UserImageModel
	if err := ctx.db.
		Select("etag, objectName").
		Where("uid = ? AND is_deleted = 0 AND is_avatar = 1", uid).
		First(&avatar).Error; err != nil {
		global.GVA_LOGGER.Error("文件查询失败", zap.String("err", err.Error()))
		return nil, er.JZError("查询失败")
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
		global.GVA_LOGGER.Error("查询失败", zap.String("err", tx.Error.Error()))
		return 0, tx.Error
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
		global.GVA_LOGGER.Error("事务开启失败", zap.String("err", begin.Error.Error()))
		return begin.Error
	}

	if err := begin.Create(&model.UserModel{
		Username: newUser.Username,
		Password: newUser.Password,
		Nickname: newUser.Nickname,
		Grade:    2,
		Status:   1,
	}).Error; err != nil {
		begin.Rollback()
		global.GVA_LOGGER.Error("创建失败", zap.String("err", err.Error()))
		return err
	}

	if err := begin.Commit().Error; err != nil {
		global.GVA_LOGGER.Error("事务提交失败", zap.String("err", err.Error()))
		return err
	}

	return nil

}

func (ctx *UserRepo) GetUserPagingRepo() (*model.UserModel, int64, error) {

	var count int64
	var user model.UserModel

	// 查询用户数量
	if err := ctx.db.Model(&model.UserModel{}).Count(&count).Error; err != nil {
		global.GVA_LOGGER.Error("查询失败", zap.String("err", err.Error()))
		return nil, 0, err
	}

	// 查询用户信息
	return &user, count, nil
}

func (ctx *UserRepo) PutUserImageRepo(uid uint64, fileType string, objectInfo *minio.UploadInfo, isAvatar bool) error {
	begin := ctx.db.Begin()
	if begin.Error != nil {
		global.GVA_LOGGER.Error("事务开启失败", zap.String("err", begin.Error.Error()))
		return begin.Error
	}

	if err := begin.Create(&model.UserImageModel{
		UID:        uid,
		ObjectName: objectInfo.Key,
		Bucket:     global.GVA_CONFIG.Minio.BucketName,
		ETag:       objectInfo.ETag,
		FileType:   fileType,
		IsAvatar:   isAvatar,
		FileSize:   uint64(objectInfo.Size),
		IsDeleted:  false,
	}).Error; err != nil {
		begin.Rollback()
		global.GVA_LOGGER.Error("创建失败", zap.String("err", err.Error()))
		return err
	}

	if err := begin.Commit().Error; err != nil {
		global.GVA_LOGGER.Error("事务提交失败", zap.String("err", err.Error()))
		return er.JZError("创建失败")
	}

	return nil
}
