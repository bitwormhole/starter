package support

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// DefaultSerivce 是默认的 CLI 服务
type DefaultSerivce struct {
	markup.Component `id:"cli-service" initMethod:"Init"`

	// public

	CLI *cli.Context `inject:"#cli-context"`

	// private

	handlerTable map[string]cli.Handler

	// chainFacade  cli.FilterChain
	// chainInner   cli.FilterChain
	chain        cli.FilterChain
	chainBuilder *cli.FilterChainBuilder
}

func (inst *DefaultSerivce) _Impl() cli.Service {
	return inst
}

// Init ...
func (inst *DefaultSerivce) Init() error {

	err := inst.initFilters()
	if err != nil {
		return err
	}

	err = inst.initHandlers()
	if err != nil {
		return err
	}

	return nil
}

func (inst *DefaultSerivce) initHandlers() error {
	list := inst.CLI.Handlers
	if list == nil {
		return nil
	}
	for _, item := range list {
		err := item.Init(inst._Impl())
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *DefaultSerivce) initFilters() error {
	list := inst.CLI.Filters
	if list == nil {
		return nil
	}
	for _, item := range list {
		err := item.Init(inst._Impl())
		if err != nil {
			return err
		}
	}
	return nil
}

// RegisterHandler ...
func (inst *DefaultSerivce) RegisterHandler(name string, h cli.Handler) error {
	if name == "" || h == nil {
		return errors.New("bad param")
	}
	table := inst.getHandlerTable()
	older := table[name]
	if older != nil {
		vlog.Warn("the command handler is duplicate, command-name: " + name)
	}
	table[name] = h
	return nil
}

// FindHandler ...
func (inst *DefaultSerivce) FindHandler(name string) (cli.Handler, error) {
	table := inst.getHandlerTable()
	h := table[name]
	if h != nil {
		return h, nil
	}
	return nil, errors.New("no command handler with name: " + name)
}

// GetHandlerNames ...
func (inst *DefaultSerivce) GetHandlerNames() []string {
	src := inst.getHandlerTable()
	dst := make([]string, 0)
	for name := range src {
		dst = append(dst, name)
	}
	return dst
}

// AddFilter ...
func (inst *DefaultSerivce) AddFilter(priority int, filter cli.Filter) error {
	if filter == nil {
		return errors.New("filter==nil")
	}
	inst.getFilterChainBuilder().Add(priority, filter)
	return nil
}

// GetFilterChain ...
func (inst *DefaultSerivce) GetFilterChain() cli.FilterChain {
	chain := inst.chain
	if chain == nil {
		builder := inst.chainBuilder
		chain = builder.Create(false)
		inst.chain = chain
	}
	return chain
}

// GetClient ...
func (inst *DefaultSerivce) GetClient(ctx context.Context) cli.Client {
	factory := inst.CLI.ClientFactory
	return factory.CreateClient(ctx)
}

// GetClientFactory ...
func (inst *DefaultSerivce) GetClientFactory() cli.ClientFactory {
	return inst.CLI.ClientFactory
}

func (inst *DefaultSerivce) getFilterChainBuilder() *cli.FilterChainBuilder {
	builder := inst.chainBuilder
	if builder == nil {
		builder = &cli.FilterChainBuilder{}
		inst.chainBuilder = builder
	}
	return builder
}

func (inst *DefaultSerivce) getHandlerTable() map[string]cli.Handler {
	table := inst.handlerTable
	if table == nil {
		table = make(map[string]cli.Handler)
		inst.handlerTable = table
	}
	return table
}

////////////////////////////////////////////////////////////////////////////////

// type defaultSerivceFilterChainFacade struct {
// 	service *DefaultSerivce
// }

// func (inst *defaultSerivceFilterChainFacade) _Impl() cli.FilterChain {
// 	return inst
// }

// func (inst *defaultSerivceFilterChainFacade) Handle(ctx *cli.TaskContext) error {
// 	target := inst.service.chainInner
// 	if target == nil {
// 		target = inst.makeInnerChain()
// 		inst.service.chainInner = target
// 	}
// 	return target.Handle(ctx)
// }

// func (inst *defaultSerivceFilterChainFacade) makeInnerChain() cli.FilterChain {
// 	builder := inst.service.getFilterChainBuilder()
// 	return builder.Create(false)
// }
