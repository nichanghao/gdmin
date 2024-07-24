package middleware

import (
	"context"
	"gitee.com/nichanghao/gdmin/common"
	"github.com/gin-gonic/gin"
)

func RequestContextHandler(params ...interface{}) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data interface{}
		var bindingMode = common.Body

		if len(params) >= 1 {
			data = params[0]
		}
		if len(params) == 2 {
			bindingMode = params[1].(common.BindingMode)
		}

		var err error
		switch bindingMode {
		case common.Query:
			err = c.ShouldBindQuery(data)
		case common.Body:
			if data != nil {
				err = c.ShouldBindJSON(data)
			}
		default:
		}

		claims, err := common.USER_CTX.GetUserClaims(c)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		// 将gin 上下文转换 go 标准库上下文
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, common.ClaimsKey, claims)
		request := common.Request{Data: data, Context: ctx}
		c.Set(common.RequestKey, &request)
		// 继续处理请求
		c.Next()
	}
}
