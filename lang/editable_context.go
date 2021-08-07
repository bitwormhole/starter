package lang

import "context"

// Context 是对 context.Context 的扩展，是个可编辑的上下文
type Context interface {
	context.Context

	SetValue(key string, value interface{})
}
