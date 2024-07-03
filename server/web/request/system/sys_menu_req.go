package system

type SysMenuReq struct {
	Title    string `json:"title" binding:"required"`    // 菜单名称
	Password string `json:"password" binding:"required"` // 密码
}
