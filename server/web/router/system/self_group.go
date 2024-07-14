package system

import (
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
)

type SelfRouter struct{}

func (*SelfRouter) InitRouter(group *gin.RouterGroup) {

	// 用户相关路由
	sysMenuGroup := group.Group("/sys/user")
	{
		sysMenuGroup.GET("/self/info", controller.SysUser.GetSelfUserInfo)
	}

}
