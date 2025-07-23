package model

import (
	"time"
)

// UserImageRelation 用户图片关联表
type UserImageModel struct {
	IID       string    `gorm:"column:IID;type:varchar(64);primaryKey;comment:图片唯一ID"`
	UID       string    `gorm:"column:UID;type:char(9);not null;index;comment:用户ID，9位数字"`
	Path      string    `gorm:"column:path;type:varchar(255);not null;comment:图片在OSS中的路径"`
	Bucket    string    `gorm:"column:bucket;type:varchar(64);not null;comment:图片所在存储桶"`
	IsAvatar  bool      `gorm:"column:is_avatar;type:tinyint(1);default:0;comment:是否为头像：1是，0不是"`
	FileSize  int       `gorm:"column:file_size;type:int;not null;comment:图片大小(字节)"`
	IsDeleted bool      `gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否逻辑删除：1是，0不是"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置表名
func (UserImageModel) TableName() string {
	return "user_image"
}
