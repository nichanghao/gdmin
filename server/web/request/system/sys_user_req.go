package system

import (
	"gitee.com/nichanghao/gdmin/common"
)

type SysUserLoginReq struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

type SysUserPageReq struct {
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Phone    string `json:"phone"`    // 手机号
	Email    string `json:"email"`    // 邮箱
	Status   uint8  `json:"status"`   // 状态
	Gender   uint8  `json:"gender"`   // 性别
	common.PageReq
}

type SysUserEditReq struct {
	Id       uint64 `json:"id" binding:"required"`     // 用户ID
	Password string `json:"password"`                  // 密码
	Nickname string `json:"nickname"`                  // 昵称
	Phone    string `json:"phone"`                     // 手机号
	Email    string `json:"email"`                     // 邮箱
	Status   uint8  `json:"status" binding:"required"` // 状态
	Gender   uint8  `json:"gender" binding:"required"` // 性别
}

type SysUserAssignRoleReq struct {
	Id      uint64   `json:"id" binding:"required"`      // 用户ID
	RoleIds []uint64 `json:"roleIds" binding:"required"` // 角色ID
}

type SysUserResetPwdReq struct {
	Id       uint64 `json:"id" binding:"required"`       // 用户ID
	Password string `json:"password" binding:"required"` // 密码
}

type SysUserAddReq struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password"`                    // 密码
	Nickname string `json:"nickname"`                    // 昵称
	Phone    string `json:"phone"`                       // 手机号
	Email    string `json:"email"`                       // 邮箱
	Status   uint8  `json:"status" binding:"required"`   // 状态
	Gender   uint8  `json:"gender" binding:"required"`   // 性别
}
