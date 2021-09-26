package filters

import (
	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

type MultilineSupportFilter struct {
	markup.Component `class:"cli-filter"`
	Priority         int `inject:"850"`
}

func (inst *MultilineSupportFilter) _Impl() cli.Filter {
	return inst
}

func (inst *MultilineSupportFilter) Init(service cli.Service) error {
	return service.AddFilter(inst.Priority, inst)
}

func (inst *MultilineSupportFilter) Handle(tc *cli.TaskContext, next cli.FilterChain) error {

	list := tc.TaskList

	if list == nil {
		current := tc.CurrentTask
		if current == nil {
			return nil
		}
		return next.Handle(tc)
	}

	for _, item := range list {
		if item == nil {
			continue
		}
		child := tc.Clone()
		child.CurrentTask = item
		err := next.Handle(child)
		if err != nil {
			return err
		}
	}

	return nil
}
