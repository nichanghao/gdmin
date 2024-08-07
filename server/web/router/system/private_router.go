package system

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/middleware"
	"gitee.com/nichanghao/gdmin/web/controller"
	"gitee.com/nichanghao/gdmin/web/request"
	"github.com/gin-gonic/gin"
)

type PrivateRouter struct{}

func (*PrivateRouter) InitRouter(group *gin.RouterGroup) {
	// 用户相关路由
	sysUserGroup := group.Group("/sys/user")
	{
		sysUserGroup.POST("page", controller.SysUser.PageUsers)
		sysUserGroup.POST("add",
			middleware.RequestContextHandler(&request.SysUserAddReq{}), controller.SysUser.AddUser)
		sysUserGroup.PUT("edit",
			middleware.RequestContextHandler(&request.SysUserEditReq{}), controller.SysUser.EditUser)
		sysUserGroup.PUT("reset-password",
			middleware.RequestContextHandler(&request.SysUserResetPwdReq{}), controller.SysUser.ResetPassword)
		sysUserGroup.DELETE("delete",
			middleware.RequestContextHandler(&request.QueryIdReq{}, common.BindModeQuery), controller.SysUser.DeleteUser)
		sysUserGroup.PUT("assign-roles",
			middleware.RequestContextHandler(&request.SysUserAssignRoleReq{}), controller.SysUser.AssignRoles)
		sysUserGroup.PUT("update-status", controller.SysUser.UpdateStatus)
	}

	// 菜单相关路由
	sysMenuGroup := group.Group("/sys/menu")
	{
		sysMenuGroup.GET("tree", controller.SysMenu.GetAllMenuTree)
		sysMenuGroup.POST("add",
			middleware.RequestContextHandler(&request.SysMenuAddReq{}), controller.SysMenu.AddMenu)
		sysMenuGroup.PUT("edit",
			middleware.RequestContextHandler(&request.SysMenuUpdateReq{}), controller.SysMenu.EditMenu)
		sysMenuGroup.DELETE("delete",
			middleware.RequestContextHandler(&request.SysMenuUpdateReq{}, common.BindModeQuery), controller.SysMenu.DeleteMenu)
	}

	// 角色相关路由
	sysRoleGroup := group.Group("/sys/role")
	{
		sysRoleGroup.POST("page", controller.SysRole.PageRoles)
		sysRoleGroup.POST("add",
			middleware.RequestContextHandler(&request.SysRoleAddReq{}), controller.SysRole.AddRole)
		sysRoleGroup.PUT("edit",
			middleware.RequestContextHandler(&request.SysRoleEditReq{}), controller.SysRole.EditRole)
		sysRoleGroup.DELETE("delete",
			middleware.RequestContextHandler(&request.QueryIdReq{}, common.BindModeQuery), controller.SysRole.DeleteRole)
		sysRoleGroup.PUT("assign-menus",
			middleware.RequestContextHandler(&request.SysAssignRoleMenuReq{}), controller.SysRole.AssignRoleMenus)

	}
}
