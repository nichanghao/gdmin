package request

import "gitee.com/nichanghao/gdmin/web/request/system"

type (
	SysUserLoginReq = system.SysUserLoginReq
	SysUserPageReq  = system.SysUserPageReq
	SysUserEditReq  = system.SysUserEditReq
	SysUserAddReq   = system.SysUserAddReq
)

type (
	SysMenuAddReq = system.SysMenuAddReq

	SysMenuUpdateReq = system.SysMenuUpdateReq
)

type (
	SysRolePageReq       = system.SysRolePageReq
	SysRoleAddReq        = system.SysRoleAddReq
	SysRoleEditReq       = system.SysRoleEditReq
	SysAssignRoleMenuReq = system.SysAssignRoleMenuReq
)

type QueryIdReq struct {
	Id uint64 `form:"id" binding:"required"`
}
