package cli

// Filter 用来过滤请求
type Filter interface {
	Init(service Service) error
	Handle(ctx *TaskContext, next FilterChain) error
}

// FilterChain 代表过滤器链条
type FilterChain interface {
	Handle(ctx *TaskContext) error
}
