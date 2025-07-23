package model

import (
	"time"
)

// SysMenu 系统菜单表
type MenuModel struct {
	ID         string    `gorm:"column:id;type:varchar(64);primaryKey;comment:菜单唯一ID"`
	MenuName   string    `gorm:"column:menu_name;type:varchar(50);not null;comment:菜单名称"`
	MenuNameCN string    `gorm:"column:menu_name_cn;type:varchar(50);comment:菜单中文名称"`
	MenuLevel  int8      `gorm:"column:menu_level;type:tinyint(1);not null;comment:菜单等级：1为主菜单，2为子菜单"`
	MenuType   string    `gorm:"column:menu_type;type:varchar(20);comment:菜单类型"`
	MenuStatus bool      `gorm:"column:menu_status;type:tinyint(1);default:1;comment:菜单状态：1启用，0禁用"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置表名
func (MenuModel) TableName() string {
	return "menus"
}
