package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct

type creationContextFacade struct {
	// impl CreationContext & RuntimeContext
	core *creationContextCore
}

type creationComponentsFacade struct {
	core *creationContextCore
}

////////////////////////////////////////////////////////////////////////////////
// impl creationContextFacade

func (inst *creationContextFacade) getParentRC() application.RuntimeContext {
	return inst.core.parent.context
}

func (inst *creationContextFacade) GetURI() string {
	return inst.getParentRC().GetURI()
}

func (inst *creationContextFacade) GetApplicationName() string {
	return inst.getParentRC().GetApplicationName()
}

func (inst *creationContextFacade) GetApplicationVersion() string {
	return inst.getParentRC().GetApplicationVersion()
}

func (inst *creationContextFacade) GetStartupTimestamp() int64 {
	return inst.getParentRC().GetStartupTimestamp()
}

func (inst *creationContextFacade) GetShutdownTimestamp() int64 {
	return inst.getParentRC().GetShutdownTimestamp()
}

func (inst *creationContextFacade) NewChild() application.RuntimeContext {
	return inst.getParentRC().NewChild()
}

func (inst *creationContextFacade) GetComponents() application.Components {
	return inst.core.components
}

func (inst *creationContextFacade) NewGetter(ec lang.ErrorCollector) application.ContextGetter {
	ctx := inst.core.proxy
	getter := &innerContextGetter{}
	getter.init(ctx, ec)
	return getter
}

func (inst *creationContextFacade) GetReleasePool() collection.ReleasePool {
	return inst.core.pool
}

func (inst *creationContextFacade) GetArguments() collection.Arguments {
	return inst.getParentRC().GetArguments()
}

func (inst *creationContextFacade) GetAttributes() collection.Attributes {
	return inst.getParentRC().GetAttributes()
}

func (inst *creationContextFacade) GetEnvironment() collection.Environment {
	return inst.getParentRC().GetEnvironment()
}

func (inst *creationContextFacade) GetProperties() collection.Properties {
	return inst.getParentRC().GetProperties()
}

func (inst *creationContextFacade) GetParameters() collection.Parameters {
	return inst.getParentRC().GetParameters()
}

func (inst *creationContextFacade) GetResources() collection.Resources {
	return inst.getParentRC().GetResources()
}

func (inst *creationContextFacade) OpenCreationContext(scope application.ComponentScope) application.CreationContext {
	return inst
}

func (inst *creationContextFacade) GetErrorHandler() lang.ErrorHandler {
	return inst.getParentRC().GetErrorHandler()
}

func (inst *creationContextFacade) SetErrorHandler(h lang.ErrorHandler) {
	inst.getParentRC().SetErrorHandler(h)
}

// as CreationContext

func (inst *creationContextFacade) GetScope() application.ComponentScope {
	return inst.core.scope
}

func (inst *creationContextFacade) GetContext() application.RuntimeContext {
	return inst.core.proxy
}

func (inst *creationContextFacade) Close() error {
	return inst.core.loader.startAllComponents()
}

////////////////////////////////////////////////////////////////////////////////
// impl creationComponentsFacade

func (inst *creationComponentsFacade) GetComponentNameList(include_aliases bool) []string {
	return inst.core.finder.listIds(include_aliases)
}

func (inst *creationComponentsFacade) GetComponent(name string) (lang.Object, error) {
	holder, err := inst.core.finder.findHolderById(name)
	if err != nil {
		return nil, err
	}
	return inst.core.loader.loadComponent(holder)
}

func (inst *creationComponentsFacade) GetComponentByClass(classSelector string) (lang.Object, error) {
	holder, err := inst.core.finder.findHolderByTypeName(classSelector)
	if err != nil {
		return nil, err
	}
	return inst.core.loader.loadComponent(holder)
}

func (inst *creationComponentsFacade) GetComponentsByClass(classSelector string) []lang.Object {
	holders := inst.core.finder.selectHoldersByTypeName(classSelector)
	instances, err := inst.core.loader.loadComponents(holders)
	if err != nil {
		inst.core.proxy.current.GetErrorHandler().OnError(err)
		return make([]lang.Object, 0)
	}
	return instances
}

func (inst *creationComponentsFacade) Export(table map[string]application.ComponentHolder) map[string]application.ComponentHolder {
	return inst.core.parent.context.GetComponents().Export(table)
}

func (inst *creationComponentsFacade) Import(src map[string]application.ComponentHolder) {
	inst.core.parent.context.GetComponents().Import(src)
}

////////////////////////////////////////////////////////////////////////////////
