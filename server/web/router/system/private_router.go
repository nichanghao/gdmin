package system

import (
	"gitee.com/nichanghao/gdmin/web/controller"
	"github.com/gin-gonic/gin"
)

type PrivateRouter struct{}

func (*PrivateRouter) InitRouter(group *gin.RouterGroup) {
	// 用户相关路由
	sysUserGroup := group.Group("/sys/user")
	{
		sysUserGroup.POST("page", controller.SysUser.PageUsers)
		sysUserGroup.PUT("edit", controller.SysUser.EditUser)
		sysUserGroup.PUT("reset-password", controller.SysUser.ResetPassword)
		sysUserGroup.DELETE("delete", controller.SysUser.DeleteUser)
		sysUserGroup.PUT("allocate-roles", controller.SysUser.AllocateRoles)
		sysUserGroup.POST("add", controller.SysUser.AddUser)
		sysUserGroup.PUT("update-status", controller.SysUser.UpdateStatus)
	}

	// 菜单相关路由
	sysMenuGroup := group.Group("/sys/menu")
	{
		sysMenuGroup.GET("tree", controller.SysMenu.GetAllMenuTree)
		sysMenuGroup.POST("add", controller.SysMenu.AddMenu)
		sysMenuGroup.PUT("edit", controller.SysMenu.EditMenu)
		sysMenuGroup.DELETE("delete", controller.SysMenu.DeleteMenu)

	}

	// 角色相关路由
	sysRoleGroup := group.Group("/sys/role")
	{
		sysRoleGroup.POST("page", controller.SysRole.PageRoles)
		sysRoleGroup.POST("add", controller.SysRole.AddRole)
		sysRoleGroup.PUT("edit", controller.SysRole.EditRole)
		sysRoleGroup.DELETE("delete", controller.SysRole.DeleteRole)

	}
}
