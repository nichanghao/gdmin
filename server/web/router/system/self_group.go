package system

import (
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
)

type SelfRouter struct{}

func (*SelfRouter) InitRouter(group *gin.RouterGroup) {

	// 菜单相关路由
	sysMenuGroup := group.Group("/sys/menu")
	{
		sysMenuGroup.GET("/self/tree", controller.SysMenu.GetSelfMenuTree)
	}

}
