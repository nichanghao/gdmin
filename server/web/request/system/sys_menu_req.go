package system

type SysMenuReq struct {
	Id            uint64 `json:"id"`                            // 菜单ID
	Name          string `json:"name" binding:"required"`       // 菜单名称
	Type          int8   `json:"type" binding:"gte=1,lte=2"`    // 菜单类型(1:菜单,2:按钮)
	Permission    string `json:"permission" binding:"required"` // 权限标识
	Path          string `json:"path" `                         // 路由地址
	Component     string `json:"component" `                    // 组件路径
	ComponentName string `json:"componentName"`                 // 组件名称
	Sort          int    `json:"sort"`                          // 排序
	ParentId      uint64 `json:"parentId"`                      // 父菜单ID
}
