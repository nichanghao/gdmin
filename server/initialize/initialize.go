package initialize

import "gitee.com/nichanghao/gdmin/global"

func init() {
	// 初始化配置文件
	InitViper()

	// 初始化日志组件
	InitZap()

	// 初始化gorm
	gorm := InitGorm()
	global.GormDB = gorm

	// 初始化 casbin
	InitCasbin()

	// 初始化 redis
	InitRedis()

	//初始化gin
	engine := InitGin()
	global.GinEngine = engine
}
