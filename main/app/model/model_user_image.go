package model

import (
	"time"
)

// UserImage 用户图片信息表结构体
type UserImageModel struct {
	ID          string    `gorm:"primaryKey;column:id;type:varchar(36);not null;comment:图片唯一标识，使用UUID" json:"imageID"`
	UID         string    `gorm:"column:uid;type:varchar(36);not null;index:idx_uid;comment:关联用户表的uid" json:"uid"`
	MinioPath   string    `gorm:"column:minio_path;type:varchar(512);not null;index:idx_minio_path,length:255;comment:MinIO存储路径" json:"minioPath"`
	MinioBucket string    `gorm:"column:minio_bucket;type:varchar(128);not null;comment:MinIO存储桶名称" json:"minioBucket"`
	IsAvatar    int8      `gorm:"column:is_avatar;type:tinyint;not null;default:0;comment:是否是头像(0:不是头像,1:是头像)" json:"IsAvatar"`
	FileSize    int64     `gorm:"column:file_size;type:bigint;not null;comment:文件大小(字节)" json:"fileSize"`
	IsDeleted   int8      `gorm:"column:is_deleted;type:tinyint;not null;default:0;comment:是否删除(0:未删除,1:已删除)" json:"isDeleted"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"`
}

// TableName 指定表名
func (UserImageModel) TableName() string {
	return "user_images"
}
