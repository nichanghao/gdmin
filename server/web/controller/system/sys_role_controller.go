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

type SysRoleController struct{}

// PageRoles 角色列表
func (*SysRoleController) PageRoles(c *gin.Context) {

	var req request.SysRolePageReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
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

	var req request.SysRoleReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	var role model.SysRole
	err := copier.Copy(&role, &req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if err2 := service.SysRole.AddRole(&role); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(true, c)
	}
}

// EditRole 编辑角色
func (*SysRoleController) EditRole(c *gin.Context) {

	var req request.SysRoleReq

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	if req.Id == 0 {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}

	var role model.SysRole
	if err := copier.Copy(&role, &req); err != nil {
		_ = c.Error(err)
		return
	}

	if err := service.SysRole.EditRole(&role); err != nil {
		_ = c.Error(err)
	} else {
		response.OkWithData(true, c)
	}
}

// DeleteRole 删除角色
func (*SysRoleController) DeleteRole(c *gin.Context) {

	id := c.Query("id")

	if id == "" {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}

	roleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		_ = c.Error(buserr.ErrIllegalParameter)
		return
	}

	if err2 := service.SysRole.DeleteRole(roleId); err2 != nil {
		_ = c.Error(err2)
	} else {
		response.OkWithData(true, c)
	}
}
