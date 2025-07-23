package model

import (
	"time"
)

// UserMenuRelation 用户菜单关联表
type UserMenuModel struct {
	ID        uint64    `gorm:"column:id;type:bigint;primaryKey;autoIncrement;comment:自增主键"`
	Mid       string    `gorm:"column:mid;type:varchar(64);not null;comment:菜单ID"`
	Uid       string    `gorm:"column:uid;type:char(9);not null;index;comment:用户ID，9位数字"`
	Status    bool      `gorm:"column:status;type:tinyint(1);default:1;comment:是否启用：1启用，0禁用"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置表名
func (UserMenuModel) TableName() string {
	return "user_menu"
}
