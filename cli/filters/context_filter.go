package filters

import (
	"context"

	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

// ContextFilter 这个过滤器用于绑定执行命令的上下文
type ContextFilter struct {
	markup.Component `class:"cli-filter"`

	Priority int             `inject:"900"`
	Context  context.Context `inject:"context"`
	Service  cli.Service
}

func (inst *ContextFilter) _Impl() cli.Filter {
	return inst
}

// Init 初始化过滤器
func (inst *ContextFilter) Init(service cli.Service) error {

	if inst.Service == nil {
		inst.Service = service
	}

	return service.AddFilter(inst.Priority, inst)
}

// Handle 处理请求
func (inst *ContextFilter) Handle(tc *cli.TaskContext, next cli.FilterChain) error {

	if tc.Context == nil {
		tc.Context = inst.Context
	}

	if tc.Service == nil {
		tc.Service = inst.Service
	}

	// for console
	console, err := cli.GetConsole(tc.Context)
	if err != nil {
		return err
	}
	tc.Console = console

	// next
	return next.Handle(tc)
}
