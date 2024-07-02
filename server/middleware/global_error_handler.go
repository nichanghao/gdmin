package middleware

import (
	"errors"
	"gitee.com/nichanghao/gdmin/model/common"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/gin-gonic/gin"
)

// GlobalErrorHandler 定义一个全局的错误处理中间件
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			// 检查是否有错误发生
			if len(c.Errors) > 0 {
				err := c.Errors.Last()

				// 业务错误
				var busErr *common.BusinessError
				switch {
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
