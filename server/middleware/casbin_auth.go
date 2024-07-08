package middleware

import (
	"gitee.com/nichanghao/gdmin/initialize"
	"github.com/gin-gonic/gin"
)

// CasbinAuth 权限控制
func CasbinAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		permission := initialize.PermissionRouter[c.HandlerName()]

	}
}
