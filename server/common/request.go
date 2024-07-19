package common

import "context"

// BindingMode 参数绑定模式，用于gin绑定请求参数
type BindingMode int

const (
	// Body 表示请求体绑定
	Body BindingMode = iota
	// Query 表示查询参数绑定
	Query
)

// Request is the base struct for all requests.
type Request struct {
	// 请求数据
	Data interface{}

	// 请求上下文
	Context context.Context
}
