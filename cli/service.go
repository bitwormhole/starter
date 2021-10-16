package cli

import "context"

// Service 是用来处理命令的服务
type Service interface {

	// for handler

	RegisterHandler(name string, h Handler) error

	FindHandler(name string) (Handler, error)

	GetHandlerNames() []string

	// for filter

	// AddFilter 添加一个过滤器到服务中。
	// priority 是过滤器的优先顺序，数值越大，优先级越高，处理顺序越靠前。
	AddFilter(priority int, filter Filter) error

	GetFilterChain() FilterChain

	// for client

	GetClient(ctx context.Context) Client

	GetClientFactory() ClientFactory
}
