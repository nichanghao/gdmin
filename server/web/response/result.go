package response

import (
	error2 "gitee.com/nichanghao/gdmin/model/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type R struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = -1
	SUCCESS = 200
)

func Result(httpCode int, code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(httpCode, R{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, "SUCCESS", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, "SUCCESS", c)
}

func OkWithResult(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(http.StatusBadRequest, ERROR, map[string]interface{}{}, "FAILED", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusBadRequest, ERROR, map[string]interface{}{}, message, c)
}

func FailWithResult(data interface{}, message string, c *gin.Context) {
	Result(http.StatusBadRequest, ERROR, data, message, c)
}

func FailWithBusErr(busErr *error2.BusinessError, c *gin.Context) {
	Result(http.StatusBadRequest, busErr.Code, nil, busErr.Error(), c)
}
