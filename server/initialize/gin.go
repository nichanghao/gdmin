package initialize

import (
	"gitee.com/nichanghao/gdmin/web/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitGin() *gin.Engine {

	engine := gin.New()

	// 异常处理
	engine.Use(gin.Recovery())

	baseGroup := engine.Group("")
	{
		// 健康检查
		baseGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	// 初始化用户相关路由
	router.SysUser.InitRouter(baseGroup)

	return engine

}
