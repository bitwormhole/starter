package runtime

import (
	"errors"
	"io"
	"sort"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// CreationContext 是构建时的上下文
type CreationContext interface {
	io.Closer
	GetScope() application.ComponentScope
	GetContext() application.Context

	ErrorCollector() lang.ErrorCollector
	HandleError(err error)
	LastError() error

	LoadComponent(h application.ComponentHolder) (lang.Object, error)
	LoadComponents(list []application.ComponentHolder) []lang.Object
}

////////////////////////////////////////////////////////////////////////////////
// struct

type contextCreation struct {
	parent            *contextRuntime
	proxy             *contextProxy
	components        application.Components
	pool              lang.ReleasePool
	scope             application.ComponentScope
	cache             map[string]*componentLoading
	errors            lang.DefaultErrorCollector
	loadingIndexCount int
}

func (inst *contextCreation) init(parent *contextRuntime, scope application.ComponentScope) CreationContext {

	proxy := &contextProxy{}
	cache := make(map[string]*componentLoading)
	pool := parent.GetReleasePool()

	if scope == 0 {
		scope = application.ScopePrototype
	}
	if scope == application.ScopePrototype {
		pool = lang.CreateReleasePool()
	}

	inst.parent = parent
	inst.proxy = proxy
	inst.cache = cache
	inst.scope = scope
	inst.pool = pool
	inst.components = parent.GetComponents()

	proxy.runtime = parent
	proxy.creation = inst
	proxy.current = inst
	proxy.pool = pool

	return inst
}

func (inst *contextCreation) getParentContext() application.Context {
	return inst.parent
}

func (inst *contextCreation) HandleError(err error) {
	inst.errors.Append(err)
}

func (inst *contextCreation) ErrorCollector() lang.ErrorCollector {
	return &inst.errors
}

func (inst *contextCreation) LastError() error {
	return inst.errors.LastError()
}

func (inst *contextCreation) GetURI() string {
	return inst.getParentContext().GetURI()
}

func (inst *contextCreation) GetApplicationName() string {
	return inst.getParentContext().GetApplicationName()
}

func (inst *contextCreation) GetApplicationVersion() string {
	return inst.getParentContext().GetApplicationVersion()
}

func (inst *contextCreation) GetStartupTimestamp() int64 {
	return inst.getParentContext().GetStartupTimestamp()
}

func (inst *contextCreation) GetShutdownTimestamp() int64 {
	return inst.getParentContext().GetShutdownTimestamp()
}

func (inst *contextCreation) NewChild() application.Context {
	return inst.getParentContext().NewChild()
}

func (inst *contextCreation) GetComponents() application.Components {
	return inst.components
}

func (inst *contextCreation) GetReleasePool() lang.ReleasePool {
	return inst.pool
}

func (inst *contextCreation) GetArguments() collection.Arguments {
	return inst.getParentContext().GetArguments()
}

func (inst *contextCreation) GetAttributes() collection.Attributes {
	return inst.getParentContext().GetAttributes()
}

func (inst *contextCreation) GetEnvironment() collection.Environment {
	return inst.getParentContext().GetEnvironment()
}

func (inst *contextCreation) GetProperties() collection.Properties {
	return inst.getParentContext().GetProperties()
}

func (inst *contextCreation) GetParameters() collection.Parameters {
	return inst.getParentContext().GetParameters()
}

func (inst *contextCreation) GetResources() collection.Resources {
	return inst.getParentContext().GetResources()
}

func (inst *contextCreation) GetErrorHandler() lang.ErrorHandler {
	return inst.getParentContext().GetErrorHandler()
}

func (inst *contextCreation) SetErrorHandler(h lang.ErrorHandler) {
	inst.getParentContext().SetErrorHandler(h)
}

// as CreationContext

func (inst *contextCreation) GetScope() application.ComponentScope {
	return inst.scope
}

func (inst *contextCreation) GetContext() application.Context {
	return inst.proxy
}

func (inst *contextCreation) Close() error {
	// inject all
	for timeout := 99; timeout > 0; timeout++ {
		cnt, err := inst.injectAllLoadings()
		if err != nil {
			return err
		}
		if cnt < 1 {
			break
		}
	}
	// close proxy
	inst.proxy.start()
	// start all
	return inst.startAllLoadings()
}

func (inst *contextCreation) listAllLoadings(reverse bool) []*componentLoading {
	table := inst.cache
	list := make([]*componentLoading, 0)
	for key := range table {
		list = append(list, table[key])
	}
	sort.Sort(&componentLoadingSorter{items: list})
	if reverse {
		i1 := 0
		i2 := len(list) - 1
		for i1 < i2 {
			tmp := list[i1]
			list[i1] = list[i2]
			list[i2] = tmp
			i1++
			i2--
		}
	}
	return list
}

func (inst *contextCreation) startAllLoadings() error {
	list := inst.listAllLoadings(true)
	for index := range list {
		loading := list[index]
		if loading == nil {
			continue
		}
		err := loading.tryStart(inst.pool)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *contextCreation) injectAllLoadings() (int, error) {
	count := 0
	ctx := inst.proxy
	list := inst.listAllLoadings(false)
	for index := range list {
		loading := list[index]
		if loading == nil {
			continue
		}
		if loading.hasInjected {
			continue
		}
		err := loading.tryInject(ctx)
		if err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

func (inst *contextCreation) getLoading(h application.ComponentHolder, create bool) (*componentLoading, error) {
	if h == nil {
		return nil, errors.New("holder==nil")
	}
	id := h.GetInfo().GetID()
	loading := inst.cache[id]
	if loading != nil {
		return loading, nil
	}
	if create {
		inst.loadingIndexCount++
		// create new
		loading = &componentLoading{}
		loading.holder = h
		loading.instance = h.GetInstance()
		loading.loadingOrder = inst.loadingIndexCount
		// result
		inst.cache[id] = loading
		return loading, nil
	}
	return nil, errors.New("no component(Loading) named: " + id)
}

func (inst *contextCreation) LoadComponent(h application.ComponentHolder) (lang.Object, error) {
	loading, err := inst.getLoading(h, true)
	if err != nil {
		return nil, err
	}
	if loading == nil {
		return nil, errors.New("componentLoading==nil")
	}
	target := loading.instance.Get()
	return target, nil
}

func (inst *contextCreation) LoadComponents(list []application.ComponentHolder) []lang.Object {
	dst := make([]lang.Object, 0)
	if list == nil {
		return dst
	}
	for index := range list {
		holder := list[index]
		if holder == nil {
			continue
		}
		target, err := inst.LoadComponent(holder)
		if err != nil {
			inst.getParentContext().GetErrorHandler().OnError(err)
			continue
		}
		dst = append(dst, target)
	}
	return dst
}

func (inst *contextCreation) FindComponent(selector string) (lang.Object, error) {
	holder, err := inst.components.GetComponent(selector)
	if err != nil {
		return nil, err
	}
	return inst.LoadComponent(holder)
}

func (inst *contextCreation) FindComponents(selector string) []lang.Object {
	holders := inst.components.GetComponents(selector)
	return inst.LoadComponents(holders)
}

func (inst *contextCreation) Injector() application.Injector {
	return inst.InjectorScope(0)
}

func (inst *contextCreation) InjectorScope(scope application.ComponentScope) application.Injector {
	injector := &innerInjector{}
	injector.init(inst, false)
	return injector
}

////////////////////////////////////////////////////////////////////////////////
