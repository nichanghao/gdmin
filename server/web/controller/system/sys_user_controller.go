package system

import (
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
)

type SysUserController struct{}

// Login 用户登录
func (sysUser *SysUserController) Login(c *gin.Context) {

	var req request.SysUserLogin

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithResult(&req, "登录成功", c)

}
