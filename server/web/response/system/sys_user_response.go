package system

import "gitee.com/nichanghao/gdmin/model"

type SysUserLoginResp struct {
	Token    string         `json:"token"`    // 用户登录token
	UserInfo *model.SysUser `json:"userInfo"` // 用户信息
}
