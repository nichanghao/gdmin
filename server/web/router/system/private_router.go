package system

import (
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PrivateRouter struct{}

func (*PrivateRouter) InitRouter(group *gin.RouterGroup) {
	// 用户相关路由
	sysUserGroup := group.Group("/sys/user")
	{

		sysUserGroup.GET("list", func(c *gin.Context) {
			c.JSON(http.StatusOK, "list")
		})
	}

	// 菜单相关路由
	sysMenuGroup := group.Group("/sys/menu")
	{
		sysMenuGroup.GET("tree", controller.SysMenu.GetMenuTree)
		sysMenuGroup.POST("add", controller.SysMenu.AddMenu)
		sysMenuGroup.PUT("edit", controller.SysMenu.EditMenu)
		sysMenuGroup.DELETE("delete", controller.SysMenu.DeleteMenu)

	}
}
