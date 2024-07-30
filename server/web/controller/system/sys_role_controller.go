package system

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
)

type SysRoleController struct{}

// PageRoles 角色列表
func (*SysRoleController) PageRoles(c *gin.Context) {

	var req request.SysRolePageReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		// 允许空的请求体
		if err.Error() != "EOF" {
			_ = c.Error(err)
			return
		}
	}
	// 初始化默认值
	req.InitDefaultValue()

	if data, err2 := service.SysRole.PageRoles(&req); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(data, c)
	}
}

// AddRole 创建角色
func (*SysRoleController) AddRole(c *gin.Context) {

	_request, _ := c.Get(common.RequestKey)
	req := _request.(*common.Request)

	if err2 := service.SysRole.AddRole(req); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(true, c)
	}
}

// EditRole 编辑角色
func (*SysRoleController) EditRole(c *gin.Context) {

	_request, _ := c.Get(common.RequestKey)

	if err := service.SysRole.EditRole(_request.(*common.Request)); err != nil {
		_ = c.Error(err)
	} else {
		response.OkWithData(true, c)
	}
}

// DeleteRole 删除角色
func (*SysRoleController) DeleteRole(c *gin.Context) {

	_request, _ := c.Get(common.RequestKey)
	if err2 := service.SysRole.DeleteRole(_request.(*common.Request)); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(true, c)
	}
}

// AssignRoleMenus 分配角色菜单
func (*SysRoleController) AssignRoleMenus(c *gin.Context) {

	_request, _ := c.Get(common.RequestKey)
	if err := service.SysRole.AssignRoleMenus(_request.(*common.Request)); err != nil {
		_ = c.Error(err)
	} else {
		response.OkWithData(true, c)
	}

}

// AllSimpleRoles 获取角色列表（用户管理页面分配用户角色时展示使用）
func (*SysRoleController) AllSimpleRoles(c *gin.Context) {

	if res, err := service.SysRole.AllSimpleRoles(); err != nil {
		_ = c.Error(err)
	} else {
		response.OkWithData(res, c)
	}

}
