package system

import (
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/service"
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
