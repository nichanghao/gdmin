package system

import "gitee.com/nichanghao/gdmin/model"

type SysUserLoginResp struct {
	Token    string         `json:"token"`    // 用户登录token
	UserInfo *model.SysUser `json:"userInfo"` // 用户信息
}

type SysUserInfoResp struct {
	User        *model.SysUser   `json:"user"`        // 用户信息
	MenuTree    []*model.SysMenu `json:"menuTree"`    // 用户菜单树
	Permissions []string         `json:"permissions"` // 用户权限标识集合
}
