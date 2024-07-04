package initialize

import (
	"gitee.com/nichanghao/gdmin/middleware"
	"gitee.com/nichanghao/gdmin/web/router"
	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {

	engine := gin.New()

	engine.Use(middleware.GlobalErrorHandler())
	// 异常处理
	engine.Use(gin.Recovery())

	// 公开路由，不需要jwt鉴权和casbin权限控制
	baseGroup := engine.Group("")
	router.Base.InitRouter(baseGroup)

	// 私有路由，需要权限控制
	privateGroup := engine.Group("")
	privateGroup.Use(middleware.JwtAuth())
	privateGroup.Use(middleware.CasbinAuth())
	// 初始化路由
	router.Private.InitRouter(privateGroup)

	return engine

}
