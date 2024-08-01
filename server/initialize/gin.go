package initialize

import (
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/middleware"
	"gitee.com/nichanghao/gdmin/web/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitGin() *gin.Engine {

	gin.SetMode(global.Config.Gin.Mode)
	engine := gin.New()
	engine.Use(middleware.GlobalErrorHandler())
	engine.Use(middleware.GinZapLogger(zap.L()))
	// panic 处理
	engine.Use(middleware.GinRecovery(zap.L()))
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
