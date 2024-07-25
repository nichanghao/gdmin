package common

import "context"

// BindMode 参数绑定模式，用于gin绑定请求参数
type BindMode int

const (
	// BindModeBody 表示请求体绑定
	BindModeBody BindMode = iota
	// BindModeQuery 表示查询参数绑定
	BindModeQuery
)

// Request is the base struct for all requests.
type Request struct {
	// 请求数据
	Data interface{}

	// 请求上下文
	Context context.Context
}
