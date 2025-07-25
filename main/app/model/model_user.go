package model

import (
	"time"
)

// User 用户表模型
type UserModel struct {
	UID       uint64    `gorm:"column:uid;primaryKey;autoIncrement:true;comment:用户唯一标识，自增ID从100000001开始" json:"uid"`
	Username  string    `gorm:"column:username;type:varchar(50);not null;uniqueIndex;comment:用户名，唯一字段" json:"username"`
	Password  string    `gorm:"column:password;type:varchar(255);not null;comment:密码(存储加密后的值)" json:"password"`
	Nickname  string    `gorm:"column:nickname;type:varchar(50);not null;comment:用户昵称" json:"nickname"`
	Grade     int8      `gorm:"column:grade;type:tinyint;not null;default:2;comment:用户等级 (1: 管理员, 2: 普通用户, 3: VIP用户等)" json:"grade"`
	Status    int8      `gorm:"column:status;type:tinyint;not null;default:1;comment:账户状态 (0: 禁用, 1: 启用, 2: 待激活)" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"`
}

// TableName 设置表名
func (UserModel) TableName() string {
	return "users"
}
