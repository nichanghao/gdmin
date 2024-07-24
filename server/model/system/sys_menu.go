package system

import (
	"encoding/json"
	"gitee.com/nichanghao/gdmin/common"
)

type SysMenu struct {
	Id         uint64          `gorm:"primarykey;comment:菜单ID" json:"id"`
	Name       string          `gorm:"type:varchar(32);comment:菜单名称" json:"name"`
	Type       int8            `gorm:"type:tinyint(1);comment:菜单类型(1:目录,2:菜单,3:按钮)" json:"type"`
	Permission string          `gorm:"type:varchar(128);comment:权限标识" json:"permission"`
	Path       string          `gorm:"type:varchar(128);comment:路由地址" json:"path"`
	Component  string          `gorm:"type:varchar(256);comment:组件" json:"component"`
	ParentId   uint64          `gorm:"default:0;comment:父菜单ID" json:"parentId"`
	Status     int8            `gorm:"type:tinyint(1);default:1;comment:菜单状态(0:禁用,1:启用)" json:"status"`
	Meta       json.RawMessage `gorm:"type:json;comment:菜单元数据" json:"meta"`
	Children   []*SysMenu      `gorm:"-" json:"children,omitempty"`
	common.BaseDO
}
