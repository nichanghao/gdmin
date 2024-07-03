package system

import (
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/utils"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
)

type SysMenuController struct{}

func (*SysMenuController) GetMenuTree(c *gin.Context) {

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
