package filters

import (
	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

type NopFilter struct {
	markup.Component `class:"cli-filter"`

	Priority int `inject:"0"`
}

func (inst *NopFilter) _Impl() cli.Filter {
	return inst
}

func (inst *NopFilter) Init(service cli.Service) error {
	return service.AddFilter(inst.Priority, inst)
}

func (inst *NopFilter) Handle(ctx *cli.TaskContext, next cli.FilterChain) error {
	return next.Handle(ctx)
}
