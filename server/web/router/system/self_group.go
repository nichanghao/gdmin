package system

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/middleware"
	"gitee.com/nichanghao/gdmin/web/controller"
	"gitee.com/nichanghao/gdmin/web/request"
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
		sysMenuGroup.GET("/all-simple-tree", controller.SysMenu.AllSimpleMenuTree)
		sysMenuGroup.GET("/list-by-role",
			middleware.RequestContextHandler(&request.QueryIdReq{}, common.BindModeQuery), controller.SysMenu.ListMenusByRoleId)
		sysMenuGroup.GET("/self/permission-routers", controller.SysMenu.GetSelfPermissionRouters)
	}

	// 角色相关路由
	sysRoleGroup := group.Group("/sys/role")
	{
		sysRoleGroup.GET("/list-by-role", controller.SysMenu.ListMenusByRoleId)
	}

}
