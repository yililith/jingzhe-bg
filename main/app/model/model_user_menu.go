package model

import "time"

// UserMenu 用户菜单关联表
type UserMenuModel struct {
	MenuID     string    `gorm:"column:menu_id;primaryKey;type:varchar(36)" json:"menuId"` // 菜单ID(UUID)
	UID        string    `gorm:"column:uid;primaryKey;type:varchar(36)" json:"uid"`        // 用户ID(UUID)
	Status     int8      `gorm:"column:status;default:1" json:"status"`                    // 状态：0-禁用，1-启用
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`      // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`      // 更新时间
}

// TableName 设置表名
func (UserMenuModel) TableName() string {
	return "user_menu"
}
