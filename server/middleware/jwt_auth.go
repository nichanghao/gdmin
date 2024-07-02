package middleware

import (
	"gitee.com/nichanghao/gdmin/model/common"
	"gitee.com/nichanghao/gdmin/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

// JwtAuth jwt token authentication middleware
func JwtAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			_ = c.Error(common.ErrTokenIllegal)
			c.Abort()
			return
		}

		arr := strings.SplitN(authorization, " ", 2)
		if len(arr) != 2 || arr[0] != "Bearer" {
			_ = c.Error(common.ErrTokenIllegal)
			c.Abort()
			return
		}

		if userClaims, err := utils.JWT.ValidateToken(arr[1]); err != nil {
			_ = c.Error(err)
			c.Abort()
		} else {
			c.Set("claims", userClaims)
			c.Next()
		}

	}
}
