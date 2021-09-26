package filters

import (
	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
)

// ContextFilter 这个过滤器用于绑定执行命令的上下文
type ContextFilter struct {
	markup.Component `class:"cli-filter"`

	Priority int          `inject:"900"`
	Context  lang.Context `inject:"context"`
	Service  cli.Service
}

func (inst *ContextFilter) _Impl() cli.Filter {
	return inst
}

func (inst *ContextFilter) Init(service cli.Service) error {

	if inst.Service == nil {
		inst.Service = service
	}

	return service.AddFilter(inst.Priority, inst)
}

func (inst *ContextFilter) Handle(tc *cli.TaskContext, next cli.FilterChain) error {

	if tc.Context == nil {
		tc.Context = inst.Context
	}

	if tc.Service == nil {
		tc.Service = inst.Service
	}

	return next.Handle(tc)
}
