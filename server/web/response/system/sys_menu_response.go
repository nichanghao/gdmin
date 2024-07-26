package system

import (
	"encoding/json"
)

type SysMenuSimpleResp struct {
	Id       string `json:"id"`       // 菜单id
	Name     string `json:"name"`     // 菜单名称
	Type     int8   `json:"type"`     // 菜单类型
	ParentId uint64 `json:"parentId"` // 父菜单id
}

type SysRoutesResp struct {
	Id         uint64           `json:"-"`         // 菜单id
	Name       string           `json:"name"`      // 路由名称
	Type       int8             `json:"-"`         // 菜单类型(1:目录,2:菜单,3:按钮)
	Path       string           `json:"path"`      // 路由地址
	Component  string           `json:"component"` // 路由组件
	ParentId   uint64           `json:"-"`         // 父菜单ID
	Permission string           `json:"-"`         // 权限标识
	Meta       json.RawMessage  `json:"meta"`      // 路由元数据
	Children   []*SysRoutesResp `gorm:"-" json:"children,omitempty"`
}

// SysPermissionRoutersResp 用户权限路由
type SysPermissionRoutersResp struct {
	Permissions []string         `json:"permissions"` // 权限列表
	Routes      []*SysRoutesResp `json:"routes"`      // 路由列表
	Home        string           `json:"home"`        // 首页路由，用户登录后默认跳转的路由
}
