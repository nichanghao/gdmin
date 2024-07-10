package initialize

import (
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
)

func init() {
	// 初始化权限路由
	{
		addPermissionRouter(controller.SysMenu.AddMenu, "sys:menu:add")
		addPermissionRouter(controller.SysMenu.EditMenu, "sys:menu:edit")
		addPermissionRouter(controller.SysMenu.DeleteMenu, "sys:menu:delete")
	}

}

func addPermissionRouter(fn gin.HandlerFunc, permission string) {
	funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	global.PermissionRouter[funcName] = permission
}
