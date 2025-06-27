package model

import "time"

type MenuModel struct {
	MenuID     string    `gorm:"column:menu_id;primaryKey;type:varchar(32);comment:菜单ID，主键"`
	MenuName   string    `gorm:"column:menu_name;type:varchar(50);not null;unique;comment:菜单英文名称"`
	MenuNameCN string    `gorm:"column:menu_name_cn;type:varchar(50);not null;comment:菜单中文名称"`
	MenuLevel  int8      `gorm:"column:menu_level;type:tinyint;not null;comment:菜单等级(1:主菜单,2:子菜单)"`
	MenuType   string    `gorm:"column:menu_type;type:varchar(50);not null;comment:菜单类型(主菜单为自身英文名，子菜单为所属主菜单英文名)"`
	MenuStatus int8      `gorm:"column:menu_status;type:tinyint;default:1;comment:菜单状态(0:禁用,1:启用)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
}

// TableName 设置表名
func (MenuModel) TableName() string {
	return "menu"
}
