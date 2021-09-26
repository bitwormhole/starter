package filters

import (
	"errors"

	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

type ExecutorFilter struct {
	markup.Component `class:"cli-filter"`

	Priority int `inject:"700"`
}

func (inst *ExecutorFilter) _Impl() cli.Filter {
	return inst
}

func (inst *ExecutorFilter) Init(service cli.Service) error {
	return service.AddFilter(inst.Priority, inst)
}

func (inst *ExecutorFilter) Handle(ctx *cli.TaskContext, next cli.FilterChain) error {

	h := ctx.Handler
	if h == nil {
		return errors.New("handler==nil")
	}

	err := h.Handle(ctx)
	if err != nil {
		return err
	}

	return next.Handle(ctx)
}
