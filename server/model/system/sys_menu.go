package system

import "gitee.com/nichanghao/gdmin/model/common"

type SysMenu struct {
	Id        uint64 `gorm:"primarykey;comment:菜单ID"`
	Title     string `gorm:"type:varchar(32);comment:菜单名称"`
	Type      int8   `gorm:"type:tinyint(1);comment:菜单类型(1:菜单,2:按钮)"`
	Component string `gorm:"type:varchar(64);comment:前端组件名称"`
	Sort      int    `gorm:"default:0;comment:菜单排序"`
	ParentId  uint64 `gorm:"default:0;comment:父菜单ID"`
	common.BaseDO
}
