package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct

type runtimeContextFacade struct {
	core *runtimeContextCore
}

type runtimeComponentsFacade struct {
	core *runtimeContextCore
}

////////////////////////////////////////////////////////////////////////////////
// impl runtimeContextFacade

func (inst *runtimeContextFacade) GetComponents() application.Components {
	return inst.core.components
}

func (inst *runtimeContextFacade) GetReleasePool() collection.ReleasePool {
	return inst.core.releasePool
}

func (inst *runtimeContextFacade) GetArguments() collection.Arguments {
	return inst.core.arguments
}

func (inst *runtimeContextFacade) GetAttributes() collection.Attributes {
	return inst.core.attributes
}

func (inst *runtimeContextFacade) GetEnvironment() collection.Environment {
	return inst.core.environment
}

func (inst *runtimeContextFacade) GetProperties() collection.Properties {
	return inst.core.properties
}

func (inst *runtimeContextFacade) GetParameters() collection.Parameters {
	return inst.core.parameters
}

func (inst *runtimeContextFacade) GetResources() collection.Resources {
	return inst.core.resources
}

func (inst *runtimeContextFacade) GetApplicationName() string {
	return inst.core.appName
}

func (inst *runtimeContextFacade) GetApplicationVersion() string {
	return inst.core.appVersion
}

func (inst *runtimeContextFacade) GetStartupTimestamp() int64 {
	return inst.core.time1
}

func (inst *runtimeContextFacade) GetShutdownTimestamp() int64 {
	return inst.core.time2
}

func (inst *runtimeContextFacade) GetURI() string {
	return inst.core.uri
}

func (inst *runtimeContextFacade) GetErrorHandler() lang.ErrorHandler {
	h := inst.core.errorHandler
	if h == nil {
		h = lang.DefaultErrorHandler()
	}
	return h
}

func (inst *runtimeContextFacade) SetErrorHandler(h lang.ErrorHandler) {
	inst.core.errorHandler = h
}

func (inst *runtimeContextFacade) OpenCreationContext(scope application.ComponentScope) application.CreationContext {
	ccc := createCreationContextCore(inst.core)
	ccc.scope = scope
	if scope == application.ScopeSingleton {
		ccc.pool = inst.core.releasePool
	} else {
		ccc.pool = collection.CreateReleasePool()
	}
	return ccc.facade
}

func (inst *runtimeContextFacade) NewChild() application.RuntimeContext {
	ctx, _ := CreateRuntimeContext(inst)
	return ctx
}

////////////////////////////////////////////////////////////////////////////////
// impl runtimeComponentsFacade

func (inst *runtimeComponentsFacade) Import(src map[string]application.ComponentHolder) {
	dst := inst.core.componentTable
	if dst == nil {
		dst = make(map[string]application.ComponentHolder)
	}
	if src == nil {
		return
	}
	ctx := inst.core.context
	for key := range src {
		holder := src[key]
		if holder == nil {
			continue
		}
		dst[key] = holder.MakeChild(ctx)
	}
	inst.core.componentTable = dst
}

func (inst *runtimeComponentsFacade) Export(dst map[string]application.ComponentHolder) map[string]application.ComponentHolder {
	src := inst.core.componentTable
	if dst == nil {
		dst = make(map[string]application.ComponentHolder)
	}
	if src == nil {
		return dst
	}
	for key := range src {
		dst[key] = src[key]
	}
	return dst
}

func (inst *runtimeComponentsFacade) GetComponentNameList(includeAliases bool) []string {
	return inst.core.finder.listIds(includeAliases)
}

func (inst *runtimeComponentsFacade) GetComponent(name string) (lang.Object, error) {
	holder, err := inst.core.finder.findHolderById(name)
	if err != nil {
		return nil, err
	}
	return inst.core.loader.loadComponent(holder)
}

func (inst *runtimeComponentsFacade) GetComponentByClass(classSelector string) (lang.Object, error) {
	holder, err := inst.core.finder.findHolderByTypeName(classSelector)
	if err != nil {
		return nil, err
	}
	return inst.core.loader.loadComponent(holder)
}

func (inst *runtimeComponentsFacade) GetComponentsByClass(classSelector string) []lang.Object {
	holders := inst.core.finder.selectHoldersByTypeName(classSelector)
	results, err := inst.core.loader.loadComponents(holders)
	if err != nil {
		inst.core.context.GetErrorHandler().OnError(err)
		return make([]lang.Object, 0)
	}
	return results
}

////////////////////////////////////////////////////////////////////////////////
