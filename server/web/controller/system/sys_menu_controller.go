package system

import (
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/model/common"
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/utils"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
)

type SysMenuController struct{}

// GetOwnerMenuTree 获取菜单树
func (*SysMenuController) GetOwnerMenuTree(c *gin.Context) {

	claims, err := utils.CLAIMS.GetUserClaims(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if menuTrees, err2 := service.SysMenu.GetMenuTreeByUserId(claims.ID); err2 != nil {
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
		_ = c.Error(common.ErrIllegalParameter)
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
		_ = c.Error(common.ErrIllegalParameter)
		return
	}
	menuId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		_ = c.Error(common.ErrIllegalParameter)
		return
	}

	if err2 := service.SysMenu.DeleteMenu(menuId); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(id, c)
	}

}
