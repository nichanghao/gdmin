package system

import (
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
)

type SelfRouter struct{}

func (*SelfRouter) InitRouter(group *gin.RouterGroup) {

	// 用户相关路由
	sysUserGroup := group.Group("/sys/user")
	{
		sysUserGroup.GET("/self/info", controller.SysUser.GetSelfUserInfo)
	}

	// 菜单相关路由
	sysMenuGroup := group.Group("/sys/menu")
	{
		sysMenuGroup.GET("/all/simple", controller.SysMenu.ListAllMenuSimple)
		sysMenuGroup.GET("/list-by-role", controller.SysMenu.ListMenusByRoleId)
	}

}
