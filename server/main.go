package main

import (
	"gitee.com/nichanghao/gdmin/global"
	_ "gitee.com/nichanghao/gdmin/initialize"
	"gitee.com/nichanghao/gdmin/web"
)

func main() {
	// 启动web服务
	web.StartServer(global.GinEngine)
}
