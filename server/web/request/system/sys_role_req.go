package system

import (
	"gitee.com/nichanghao/gdmin/common"
)

type SysRolePageReq struct {
	Name           string `json:"name"`   // 名称查询
	Code           string `json:"code"`   // code查询
	Status         int8   `json:"status"` // 状态(1:启用 2:禁用)
	common.PageReq        // 分页数据
}

type SysRoleAddReq struct {
	Name   string `json:"name" binding:"required"` // 名称
	Code   string `json:"code" binding:"required"` // code
	Status int8   `json:"status"`                  // 状态(1:启用 2:禁用)
	Desc   string `json:"desc"`                    // 描述
}

type SysRoleEditReq struct {
	Id uint64 `json:"id" binding:"required"` // ID
	SysRoleAddReq
}

type SysAssignRoleMenuReq struct {
	RoleId  uint64   `json:"roleId" binding:"required"`  // 角色id
	MenuIds []uint64 `json:"menuIds" binding:"required"` // 菜单id集合
}
