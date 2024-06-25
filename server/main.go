package main

import (
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/initialize"
	"gitee.com/nichanghao/gdmin/web"
)

func main() {

	// 初始化配置文件
	initialize.InitViper()

	// 初始化 redis
	initialize.InitRedis()

	// 初始化gin
	engine := initialize.InitGin()

	// 初始化gorm
	gorm := initialize.InitGorm()
	global.GormDB = gorm

	// 启动web服务
	web.StartServer(engine)

}
