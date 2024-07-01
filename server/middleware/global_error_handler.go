package middleware

import (
	"errors"
	error2 "gitee.com/nichanghao/gdmin/model/common"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// GlobalErrorHandler 定义一个全局的错误处理中间件
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			// 检查是否有错误发生
			if len(c.Errors) > 0 {
				err := c.Errors.Last()

				// 业务错误
				var busErr *error2.BusinessError
				switch {
				case errors.As(err.Err, &jwt.ErrTokenExpired):
					response.Result(http.StatusBadRequest, -2, nil, "Token expired", c)
				case errors.As(err.Err, &busErr):
					response.FailWithBusErr(busErr, c)
				default:
					response.FailWithMessage(err.Error(), c)
				}

				// 中止请求
				c.Abort()
			}

		}()

		c.Next()
	}
}
