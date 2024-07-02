package system

import (
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct{}

func (*BaseRouter) InitRouter(group *gin.RouterGroup) {

	// 健康检查
	group.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// 登录
	group.POST("/login", controller.SysUser.Login)

}
