package common

// token相关异常
var (
	ErrTokenExpired = NewBusErr(10001, "token已过期！")

	ErrTokenIllegal = NewBusErr(10002, "illegal token！")
)

var (
	ErrPassWdNonMatched = NewBusErr(20001, "用户名或密码错误！")
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

func (e *BusinessError) Error() string {
	return e.Message
}
