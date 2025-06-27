package model

import (
	"time"
)

type UserModel struct {
	UID       string    `gorm:"primaryKey;column:uid;type:varchar(36);not null" json:"uid"`                                                       // 用户唯一标识，使用UUID
	Username  string    `gorm:"uniqueIndex;column:username;type:varchar(50);not null" json:"username"`                                            // 用户名，唯一字段
	Password  string    `gorm:"column:password;type:varchar(255);not null" json:"password"`                                                       // 密码(存储加密后的值)
	Nickname  string    `gorm:"column:nickname;type:varchar(50);not null" json:"nickname"`                                                        // 用户昵称
	Grade     int8      `gorm:"column:grade;type:tinyint;default:2;not null" json:"grade"`                                                        // 用户等级 (1: 管理员, 2: 普通用户, 3: VIP用户等)
	Status    int8      `gorm:"index;column:status;type:tinyint;default:1;not null" json:"status"`                                                // 账户状态 (0: 禁用, 1: 启用, 2: 待激活) 	// 最后登录时间
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"createdAt"`                             // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null" json:"updatedAt"` // 更新时间
}

// TableName 指定表名
func (UserModel) TableName() string {
	return "users"
}
