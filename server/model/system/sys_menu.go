package system

import (
	"gitee.com/nichanghao/gdmin/common"
)

type SysMenu struct {
	Id         uint64     `gorm:"primarykey;comment:菜单ID" json:"id"`
	Title      string     `gorm:"type:varchar(32);comment:菜单名称" json:"title"`
	Type       int8       `gorm:"type:tinyint(1);comment:菜单类型(1:菜单,2:按钮)" json:"type"`
	Permission string     `gorm:"type:varchar(128);comment:权限标识" json:"permission"`
	Component  string     `gorm:"type:varchar(64);comment:前端组件名称" json:"component"`
	Sort       int        `gorm:"default:0;comment:菜单排序" json:"sort"`
	ParentId   uint64     `gorm:"default:0;comment:父菜单ID" json:"parentId"`
	Children   []*SysMenu `gorm:"-" json:"children"`
	common.BaseDO
}
