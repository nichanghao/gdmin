package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// CasbinAuth 权限控制
func CasbinAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		fmt.Println(c.HandlerName())
	}
}
