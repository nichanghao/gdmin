package system

type SysMenuReq struct {
	Id        uint64 `json:"id"`                           // 菜单ID
	Title     string `json:"title" binding:"required"`     // 菜单名称
	Resource  string `json:"resource"`                     // 请求后端资源
	Action    string `json:"action"`                       // 请求资源方式
	Type      int8   `json:"type" binding:"gte=1,lte=2"`   // 菜单类型(1:菜单,2:按钮)
	Component string `json:"component" binding:"required"` // 前端组件名称
	Sort      int    `json:"sort"`                         // 排序
	ParentId  uint64 `json:"parentId"`                     // 父菜单ID
}
