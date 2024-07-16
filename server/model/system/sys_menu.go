package system

import (
	"gitee.com/nichanghao/gdmin/common"
)

type SysMenu struct {
	Id            uint64     `gorm:"primarykey;comment:菜单ID" json:"id"`
	Name          string     `gorm:"type:varchar(32);comment:菜单名称" json:"name"`
	Type          int8       `gorm:"type:tinyint(1);comment:菜单类型(1:菜单,2:按钮)" json:"type"`
	Permission    string     `gorm:"type:varchar(128);comment:权限标识" json:"permission"`
	Path          string     `gorm:"type:varchar(128);comment:路由地址" json:"path"`
	Component     string     `gorm:"type:varchar(256);comment:组件路径" json:"component"`
	ComponentName string     `gorm:"type:varchar(64);comment:组件名称" json:"componentName"`
	Sort          int        `gorm:"default:0;comment:菜单排序" json:"sort"`
	ParentId      uint64     `gorm:"default:0;comment:父菜单ID" json:"parentId"`
	Children      []*SysMenu `gorm:"-" json:"children,omitempty"`
	common.BaseDO
}
