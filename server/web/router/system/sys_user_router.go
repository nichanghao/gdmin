package system

import (
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
)

type SysUserRouter struct{}

func (user *SysUserRouter) InitRouter(group *gin.RouterGroup) {
	sysUserGroup := group.Group("/sys/user")
	{

		sysUserGroup.POST("login", controller.SysUser.Login)
	}
}
