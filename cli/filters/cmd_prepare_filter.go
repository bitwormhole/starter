package filters

import (
	"errors"

	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

// HandlerFinderFilter 这个过滤器用于绑定执行命令的handler
type HandlerFinderFilter struct {
	markup.Component `class:"cli-filter"`

	Priority int `inject:"800"`
}

func (inst *HandlerFinderFilter) _Impl() cli.Filter {
	return inst
}

func (inst *HandlerFinderFilter) Init(service cli.Service) error {
	return service.AddFilter(inst.Priority, inst)
}

func (inst *HandlerFinderFilter) Handle(ctx *cli.TaskContext, next cli.FilterChain) error {

	task := ctx.CurrentTask
	service := ctx.Service

	if task == nil || service == nil {
		return errors.New(" task == nil || service == nil ")
	}

	name := task.CommandName
	if name == "" {
		name = inst.getCommandName(task.Arguments)
		task.CommandName = name
	}

	h, err := service.FindHandler(name)
	if err != nil {
		return err
	}
	ctx.Handler = h

	return next.Handle(ctx)
}

func (inst *HandlerFinderFilter) getCommandName(args []string) string {
	if args == nil {
		return ""
	}
	if len(args) < 1 {
		return ""
	}
	return args[0]
}
