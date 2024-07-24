package middleware

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/service"
	"github.com/gin-gonic/gin"
)

// CasbinAuthHandler 权限控制
func CasbinAuthHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		userClaims, err := common.USER_CTX.GetUserClaims(c)
		// jwt token 验证失败
		if err != nil || userClaims == nil || userClaims.ID == 0 {
			_ = c.Error(buserr.ErrPermissionDenied)
			c.Abort()
			return
		}

		// 验证权限
		permission, exist := global.PermissionRouter[c.HandlerName()]
		if !exist {
			_ = c.Error(buserr.ErrPermissionDenied)
			c.Abort()
			return
		}
		if enforce, err2 := global.Enforcer.Enforce(service.SysCasbin.GetCasbinUserStr(userClaims.ID), permission); err2 != nil || !enforce {
			_ = c.Error(buserr.ErrPermissionDenied)
			c.Abort()
		} else {
			c.Next()
		}

	}
}
