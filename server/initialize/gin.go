package initialize

import (
	"gitee.com/nichanghao/gdmin/middleware"
	"gitee.com/nichanghao/gdmin/web/router"
	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {

	engine := gin.New()

	engine.Use(middleware.GlobalErrorHandler())
	// panic 处理
	engine.Use(gin.Recovery())
	// 跨域
	engine.Use(middleware.CorsHandler())

	// 公开路由组，不需要jwt鉴权和casbin权限控制
	baseGroup := engine.Group("")
	router.Base.InitRouter(baseGroup)

	//自有路由组，只要jwt鉴权即可
	selfGroup := engine.Group("")
	selfGroup.Use(middleware.JwtAuthHandler())
	router.Self.InitRouter(selfGroup)

	// 私有路由组，需要jwt鉴权和casbin权限控制
	privateGroup := engine.Group("")
	privateGroup.Use(middleware.JwtAuthHandler())
	privateGroup.Use(middleware.CasbinAuthHandler())
	router.Private.InitRouter(privateGroup)

	return engine

}
