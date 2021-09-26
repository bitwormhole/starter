package cli

import (
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// DefaultSerivce 是默认的 CLI 服务
type DefaultSerivce struct {
	markup.Component `id:"cli-service" initMethod:"Init"`

	// public

	Filters  []Filter  `inject:"cli-filter"`
	Handlers []Handler `inject:"cli-handler"`

	// private

	handlerTable map[string]Handler

	client       Client
	chainFacade  FilterChain
	chainInner   FilterChain
	chainBuilder *filterChainBuilder
}

func (inst *DefaultSerivce) _Impl() Service {
	return inst
}

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
	list := inst.Handlers
	if list == nil {
		return nil
	}
	for _, item := range list {
		err := item.Init(inst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *DefaultSerivce) initFilters() error {
	list := inst.Filters
	if list == nil {
		return nil
	}
	for _, item := range list {
		err := item.Init(inst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *DefaultSerivce) RegisterHandler(name string, h Handler) error {
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

func (inst *DefaultSerivce) FindHandler(name string) (Handler, error) {
	table := inst.getHandlerTable()
	h := table[name]
	if h != nil {
		return h, nil
	}
	return nil, errors.New("no command handler with name: " + name)
}

func (inst *DefaultSerivce) GetHandlerNames() []string {
	src := inst.getHandlerTable()
	dst := make([]string, 0)
	for name := range src {
		dst = append(dst, name)
	}
	return dst
}

func (inst *DefaultSerivce) AddFilter(priority int, filter Filter) error {
	if filter == nil {
		return errors.New("filter==nil")
	}
	inst.getFilterChainBuilder().add(priority, filter)
	return nil
}

func (inst *DefaultSerivce) GetFilterChain() FilterChain {
	chain := inst.chainFacade
	if chain == nil {
		facade := &defaultSerivceFilterChainFacade{
			service: inst,
		}
		chain = facade
		inst.chainFacade = chain
	}
	return chain
}

func (inst *DefaultSerivce) GetClient() Client {
	client := inst.client
	if client == nil {
		pNew := &DefaultClient{
			Service: inst,
		}
		client = pNew
		inst.client = client
	}
	return client
}

func (inst *DefaultSerivce) getFilterChainBuilder() *filterChainBuilder {
	builder := inst.chainBuilder
	if builder == nil {
		builder = &filterChainBuilder{}
		inst.chainBuilder = builder
	}
	return builder
}

func (inst *DefaultSerivce) getHandlerTable() map[string]Handler {
	table := inst.handlerTable
	if table == nil {
		table = make(map[string]Handler)
		inst.handlerTable = table
	}
	return table
}

////////////////////////////////////////////////////////////////////////////////

type defaultSerivceFilterChainFacade struct {
	service *DefaultSerivce
}

func (inst *defaultSerivceFilterChainFacade) _Impl() FilterChain {
	return inst
}

func (inst *defaultSerivceFilterChainFacade) Handle(ctx *TaskContext) error {
	target := inst.service.chainInner
	if target == nil {
		target = inst.makeInnerChain()
		inst.service.chainInner = target
	}
	return target.Handle(ctx)
}

func (inst *defaultSerivceFilterChainFacade) makeInnerChain() FilterChain {
	builder := inst.service.getFilterChainBuilder()
	return builder.create()
}
