package system

import (
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
)

type OwnerRouter struct{}

func (*OwnerRouter) InitRouter(group *gin.RouterGroup) {

	// 菜单相关路由
	sysMenuGroup := group.Group("/sys/menu")
	{
		sysMenuGroup.GET("/owner/tree", controller.SysMenu.GetOwnerMenuTree)
	}

}
