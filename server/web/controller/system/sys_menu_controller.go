package system

import (
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/utils"
	"github.com/gin-gonic/gin"
)

type SysMenuController struct{}

func (*SysMenuController) GetMenuTree(c *gin.Context) {

	claims, err := utils.CLAIMS.GetUserClaims(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	service.SysMenu.GetMenuTreeByUserId(claims.ID)

}
