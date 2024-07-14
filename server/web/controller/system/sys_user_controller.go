package system

import (
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/utils"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysUserController struct{}

// Login 用户登录
func (*SysUserController) Login(c *gin.Context) {

	var req request.SysUserLoginReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	sysUser := &model.SysUser{Username: req.Username, Password: req.Password}
	if resp, err := service.SysUser.Login(sysUser); err != nil {
		_ = c.Error(err)
	} else {
		response.OkWithResult(&resp, "登录成功", c)
	}

}

// GetSelfUserInfo 获取当前用户信息
func (*SysUserController) GetSelfUserInfo(c *gin.Context) {

	claims, err := utils.CLAIMS.GetUserClaims(c)
	if err != nil {
		response.FailWithMessage("获取用户信息失败！", c)
		return
	}

	if info, err2 := service.SysUser.GetSelfUserInfo(claims.ID); err2 != nil {
		response.FailWithMessage("获取用户信息失败！", c)
	} else {
		response.OkWithData(info, c)
	}
}

// PageUsers 分页查询用户列表
func (*SysUserController) PageUsers(c *gin.Context) {
	var req request.SysUserPageReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	// 初始化默认值
	req.InitDefaultValue()

	if users, err := service.SysUser.PageUsers(&req); err != nil {
		_ = c.Error(err)
	} else {
		response.OkWithData(users, c)
	}
}

// EditUser 编辑用户
func (*SysUserController) EditUser(c *gin.Context) {
	var req request.SysUserEditReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	var user model.SysUser
	if err := copier.Copy(&user, &req); err != nil {
		_ = c.Error(err)
		return
	}

	if err := service.SysUser.EditUser(&user); err != nil {
		_ = c.Error(err)
	} else {
		response.Ok(c)
	}

}

// ResetPassword 重置密码
func (*SysUserController) ResetPassword(c *gin.Context) {
	var req request.SysUserEditReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	if req.Password == "" {
		response.FailWithMessage("密码不能为空！", c)
		return
	}

	if err := service.SysUser.ResetPassword(&req); err != nil {
		_ = c.Error(err)
	} else {
		response.Ok(c)
	}

}

// DeleteUser 删除用户
func (*SysUserController) DeleteUser(c *gin.Context) {
	var req request.SysUserEditReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	if err := service.SysUser.DeleteUser(req.Id); err != nil {
		_ = c.Error(err)
	} else {
		response.Ok(c)
	}

}

// AllocateRoles 分配角色给用户
func (*SysUserController) AllocateRoles(c *gin.Context) {
	var req request.SysUserEditReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	if len(req.RoleIds) == 0 {
		response.FailWithMessage("请选择角色！", c)
		return
	}

	if err := service.SysUser.AllocateRoles(&req); err != nil {
		_ = c.Error(err)
	} else {
		response.Ok(c)
	}

}
