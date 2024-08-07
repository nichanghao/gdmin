package middleware

import (
	"gitee.com/nichanghao/gdmin/cache"
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

// JwtAuthHandler jwt token authentication middleware
func JwtAuthHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			_ = c.Error(buserr.NewTokenAuthErr("invalid token"))
			c.Abort()
			return
		}

		arr := strings.SplitN(authorization, " ", 2)
		if len(arr) != 2 || arr[0] != "Bearer" {
			_ = c.Error(buserr.NewTokenAuthErr("invalid token"))
			c.Abort()
			return
		}

		if userClaims, err := utils.JWT.ValidateToken(arr[1]); err != nil {
			_ = c.Error(err)
			c.Abort()
		} else {
			status, _ := cache.SysUserCache.GetSysUserStatus(userClaims.ID)
			if status != 1 {
				_ = c.Error(buserr.NewTokenAuthErr("用户状态异常，请联系管理员！"))
				c.Abort()
				return
			}

			c.Set(common.ClaimsKey, userClaims)
			c.Next()
		}

	}
}
