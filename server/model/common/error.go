package common

var (
	ErrIllegalParameter = NewBusErr(20001, "请求参数错误！")
)

const (

	// NoticeCode 前端统一弹框提示的code码
	NoticeCode = 201

	// TokenAuthErrCode token异常code，需要重定向至登录页面
	TokenAuthErrCode = 401
)

// BusinessError 自定义业务异常
type BusinessError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewBusErr(code int, message string) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
	}
}

func NewNoticeBusErr(message string) *BusinessError {
	return &BusinessError{
		Code:    NoticeCode,
		Message: message,
	}
}

func NewTokenAuthErr(message string) *BusinessError {
	return &BusinessError{
		Code:    TokenAuthErrCode,
		Message: message,
	}
}

func (e *BusinessError) Error() string {
	return e.Message
}
