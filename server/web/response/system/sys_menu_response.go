package system

type SysMenuSimpleResp struct {
	Id       string `json:"id"`       // 菜单id
	Name     string `json:"name"`     // 菜单名称
	Type     int8   `json:"type"`     // 菜单类型
	ParentId uint64 `json:"parentId"` // 父菜单id
}
