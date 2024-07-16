package system

import (
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
)

type SysMenuController struct{}

// GetAllMenuTree 获取所有菜单树
func (*SysMenuController) GetAllMenuTree(c *gin.Context) {

	if menuTrees, err2 := service.SysMenu.GetAllMenuTree(); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(menuTrees, c)
	}

}

// AddMenu 添加菜单
func (*SysMenuController) AddMenu(c *gin.Context) {

	var req request.SysMenuReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	var menu model.SysMenu
	err := copier.Copy(&menu, &req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if id, err2 := service.SysMenu.AddMenu(&menu); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(id, c)
	}

}

// EditMenu 编辑菜单
func (*SysMenuController) EditMenu(c *gin.Context) {

	var req request.SysMenuReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	if req.Id == 0 {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}

	var menu model.SysMenu
	err := copier.Copy(&menu, &req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if id, err2 := service.SysMenu.EditMenu(&menu); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(id, c)
	}

}

// DeleteMenu 删除菜单
func (*SysMenuController) DeleteMenu(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}
	menuId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}

	if err2 := service.SysMenu.DeleteMenu(menuId); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(id, c)
	}

}

// ListAllMenuSimple 获取所有菜单简要信息（角色管理页面分配角色权限时展示使用）
func (*SysMenuController) ListAllMenuSimple(c *gin.Context) {
	if res, err := service.SysRole.ListAllMenuSimple(); err != nil {
		response.FailWithMessage("获取菜单简要信息失败！", c)
	} else {
		response.OkWithData(res, c)
	}
}

// ListMenusByRoleId 获取角色拥有的菜单（角色管理页面分配角色权限时展示使用）
func (*SysMenuController) ListMenusByRoleId(c *gin.Context) {
	id := c.Query("roleId")
	if id == "" {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}

	roleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}

	if res, err2 := service.SysRole.ListMenusByRoleId(roleId); err2 != nil {
		response.FailWithMessage("获取角色菜单失败！", c)
	} else {
		response.OkWithData(res, c)
	}

}
