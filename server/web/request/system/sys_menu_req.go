package system

import "encoding/json"

type SysMenuAddReq struct {
	Name       string          `json:"name" binding:"required"`    // 菜单名称
	Type       int8            `json:"type" binding:"gte=1,lte=2"` // 菜单类型(1:目录,2:菜单)
	Permission string          `json:"permission"`                 // 权限标识
	Path       string          `json:"path"`                       // 路由地址
	Component  string          `json:"component"`                  // 组件
	Status     int8            `json:"status"`                     // 菜单状态(0:禁用,1:启用)
	Meta       json.RawMessage `json:"meta"`                       // 元数据
	Buttons    json.RawMessage `json:"buttons"`                    // 按钮
	ParentId   uint64          `json:"parentId"`                   // 父菜单ID
}

type SysMenuUpdateReq struct {
	Id uint64 `json:"id" form:"id" binding:"required"` // 菜单ID
	SysMenuAddReq
}
