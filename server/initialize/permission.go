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
		addPermissionRouter(controller.SysMenu.GetAllMenuTree, "sys:menu")
		addPermissionRouter(controller.SysMenu.AddMenu, "sys:menu:add")
		addPermissionRouter(controller.SysMenu.EditMenu, "sys:menu:edit")
		addPermissionRouter(controller.SysMenu.DeleteMenu, "sys:menu:delete")
	}

	{
		addPermissionRouter(controller.SysUser.PageUsers, "sys:user")
		addPermissionRouter(controller.SysUser.AddUser, "sys:user:add")
		addPermissionRouter(controller.SysUser.EditUser, "sys:user:edit")
		addPermissionRouter(controller.SysUser.ResetPassword, "sys:user:resetPwd")
		addPermissionRouter(controller.SysUser.DeleteUser, "sys:user:delete")
		addPermissionRouter(controller.SysUser.AssignRoles, "sys:user:assignRoles")
	}

	{
		addPermissionRouter(controller.SysRole.PageRoles, "sys:role")
		addPermissionRouter(controller.SysRole.AddRole, "sys:role:add")
		addPermissionRouter(controller.SysRole.EditRole, "sys:role:edit")
		addPermissionRouter(controller.SysRole.DeleteRole, "sys:role:delete")
		addPermissionRouter(controller.SysRole.AssignRoleMenus, "sys:role:assignMenus")
	}

}

func addPermissionRouter(fn gin.HandlerFunc, permission string) {
	funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	global.PermissionRouter[funcName] = permission
}
