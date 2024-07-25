package system

import "gitee.com/nichanghao/gdmin/model"

type SysMenuSimpleResp struct {
	Id       string `json:"id"`       // 菜单id
	Name     string `json:"name"`     // 菜单名称
	Type     int8   `json:"type"`     // 菜单类型
	ParentId uint64 `json:"parentId"` // 父菜单id
}

// SysPermissionRoutersResp 用户权限路由
type SysPermissionRoutersResp struct {
	Permissions []string         `json:"permissions"` // 权限列表
	Routes      []*model.SysMenu `json:"routes"`      // 路由列表
	Home        string           `json:"home"`        // 首页路由，用户登录后默认跳转的路由
}
