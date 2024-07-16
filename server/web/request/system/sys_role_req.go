package system

import (
	"gitee.com/nichanghao/gdmin/common"
)

type SysRolePageReq struct {
	Name           string `json:"name"` // 名称查询
	Code           string `json:"code"` // code查询
	common.PageReq        // 分页数据
}

type SysRoleReq struct {
	Id   uint64 `json:"id"`                      // ID
	Name string `json:"name" binding:"required"` // 名称
	Code string `json:"code" binding:"required"` // code
	Desc string `json:"desc"`                    // 描述
}

type SysAssignRoleMenuReq struct {
	RoleId  uint64   `json:"roleId" binding:"required"`  // 角色id
	MenuIds []uint64 `json:"menuIds" binding:"required"` // 描述
}
