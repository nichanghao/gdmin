package system

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
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

	claims, err := common.USER_CTX.GetUserClaims(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if info, err2 := service.SysUser.GetSelfUserInfo(claims.ID); err2 != nil {
		_ = c.Error(err)
	} else {
		response.OkWithData(info, c)
	}
}

// PageUsers 分页查询用户列表
func (*SysUserController) PageUsers(c *gin.Context) {
	var req request.SysUserPageReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil && err.Error() != "EOF" {
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

// AddUser 新增用户
func (*SysUserController) AddUser(c *gin.Context) {
	_request, _ := c.Get(common.RequestKey)
	if err := service.SysUser.AddUser(_request.(*common.Request)); err != nil {
		_ = c.Error(err)
	} else {
		response.Ok(c)
	}

}

// EditUser 编辑用户
func (*SysUserController) EditUser(c *gin.Context) {
	_request, _ := c.Get(common.RequestKey)

	if err := service.SysUser.EditUser(_request.(*common.Request)); err != nil {
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
	_request, _ := c.Get(common.RequestKey)
	if err := service.SysUser.DeleteUser(_request.(*common.Request)); err != nil {
		_ = c.Error(err)
	} else {
		response.Ok(c)
	}

}

// AssignRoles 分配角色给用户
func (*SysUserController) AssignRoles(c *gin.Context) {
	var req request.SysUserEditReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	if err := service.SysUser.AssignRoles(&req); err != nil {
		_ = c.Error(err)
	} else {
		response.Ok(c)
	}

}

// UpdateStatus 更新用户状态
func (*SysUserController) UpdateStatus(c *gin.Context) {
	var req request.SysUserEditReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	if req.Status == 0 {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}

	if err := service.SysUser.UpdateStatus(&req); err != nil {
		_ = c.Error(err)
	} else {
		response.Ok(c)
	}
}
