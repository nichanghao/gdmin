package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

// GinZapLogger 创建一个 Gin 中间件，使用 zap 记录日志
func GinZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Use CustomResponseWriter to capture response body
		crw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = crw

		// Log request information
		reqBody, _ := c.GetRawData()

		// Write back the request body
		if len(reqBody) > 0 {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		// Execute request handlers and other middleware functions
		c.Next()

		// 记录请求结束时间
		end := time.Now()
		cost := end.Sub(start)

		// 如果有错误，记录错误信息
		if len(c.Errors) > 0 {
			last := c.Errors.Last()
			logger.Error(c.Request.Method+" "+path,
				zap.Error(last),
				zap.String("query", query),
				zap.String("req", string(reqBody)))
		} else {
			// 记录请求信息
			logger.Info(c.Request.Method+" "+path,
				zap.Int("status", c.Writer.Status()),
				zap.String("req", string(reqBody)),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("resp", string(crw.body.Bytes())),
				zap.String("cost", cost.String()),
			)
		}
	}
}

// CustomResponseWriter encapsulates gin ResponseWriter to capture the response body.
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
