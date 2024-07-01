package system

import (
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
)

type SysUserController struct{}

// Login 用户登录
func (*SysUserController) Login(c *gin.Context) {

	var req request.SysUserLogin

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
