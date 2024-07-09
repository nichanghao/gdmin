package middleware

import (
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model/common"
	"gitee.com/nichanghao/gdmin/service"
	"gitee.com/nichanghao/gdmin/utils"
	"github.com/gin-gonic/gin"
)

// CasbinAuth 权限控制
func CasbinAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		userClaims, err := utils.CLAIMS.GetUserClaims(c)
		// jwt token 验证失败
		if err != nil || userClaims == nil || userClaims.ID == 0 {
			_ = c.Error(common.ErrPermissionDenied)
			c.Abort()
			return
		}

		// 验证权限
		permission, exist := global.PermissionRouter[c.HandlerName()]
		if !exist {
			_ = c.Error(common.ErrPermissionDenied)
			c.Abort()
			return
		}
		if enforce, err := global.Enforcer.Enforce(service.SysCasbin.GetCasbinUserStr(userClaims.ID), permission); err != nil || !enforce {
			_ = c.Error(common.ErrPermissionDenied)
			c.Abort()
		} else {
			c.Next()
		}

	}
}
