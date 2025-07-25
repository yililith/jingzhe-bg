package model

import (
	"time"
)

// UserImage 用户图片关联表
type UserImageModel struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UID        uint64    `gorm:"not null;index:idx_uid;column:uid" json:"uid"`
	ObjectName string    `gorm:"type:varchar(255);not null;column:objectName" json:"objectName"`                 // OSS中的对象名
	Bucket     string    `gorm:"type:varchar(64);not null;column:bucket" json:"bucket"`                          // 存储桶
	ETag       string    `gorm:"type:char(32);not null;uniqueIndex:uniq_etag;column:etag" json:"etag"`           // 文件内容MD5
	FileType   string    `gorm:"type:varchar(64);not null;column:file_type" json:"fileType"`                     // 文件类型，例如 image/webp
	IsAvatar   bool      `gorm:"not null;default:false;index:idx_is_avatar;column:is_avatar" json:"isAvatar"`    // 是否头像
	FileSize   uint64    `gorm:"not null;column:file_size" json:"fileSize"`                                      // 文件大小，字节
	IsDeleted  bool      `gorm:"not null;default:false;index:idx_is_deleted;column:is_deleted" json:"isDeleted"` // 逻辑删除标记
	CreatedAt  time.Time `gorm:"autoCreateTime;column:created_at" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updatedAt"`
}

// TableName 自定义表名
func (UserImageModel) TableName() string {
	return "user_image"
}
