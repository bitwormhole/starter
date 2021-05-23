package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct contextRuntime

type contextRuntime struct {

	// collection
	releasePool lang.ReleasePool
	components  application.Components

	arguments   collection.Arguments
	attributes  collection.Attributes
	environment collection.Environment
	properties  collection.Properties
	parameters  collection.Parameters
	resources   collection.Resources

	// helper
	errorHandler lang.ErrorHandler

	// info
	time1      int64
	time2      int64
	appName    string
	appVersion string
	uri        string
}

func (inst *contextRuntime) Init(parent application.RuntimeContext) (application.Context, error) {
	if parent == nil {
		return inst._init1()
	} else {
		return inst._init2(parent)
	}
}

func (inst *contextRuntime) _init1() (application.Context, error) {

	components := &componentTable{}
	components.context = inst
	components.table = make(map[string]application.ComponentHolder)

	inst.arguments = collection.CreateArguments()
	inst.attributes = collection.CreateAttributes()
	inst.components = components
	inst.environment = collection.CreateEnvironment()
	inst.parameters = collection.CreateParameters()
	inst.properties = collection.CreateProperties()
	inst.releasePool = lang.CreateReleasePool()
	inst.resources = nil

	inst.time1 = 0 // now

	return inst, nil
}

func (inst *contextRuntime) _init2(parent application.RuntimeContext) (application.Context, error) {

	// create new
	child := inst
	if parent == nil {
		return child, nil
	}

	// export & import
	inst.resources = parent.GetResources()
	child.GetArguments().Import(parent.GetArguments().Export())
	child.GetAttributes().Import(parent.GetAttributes().Export(nil))
	child.GetEnvironment().Import(parent.GetEnvironment().Export(nil))
	child.GetParameters().Import(parent.GetParameters().Export(nil))
	child.GetProperties().Import(parent.GetProperties().Export(nil))
	child.GetComponents().Import(parent.GetComponents().Export(nil))

	return child, nil
}

func (inst *contextRuntime) GetComponents() application.Components {
	return inst.components
}

func (inst *contextRuntime) GetReleasePool() lang.ReleasePool {
	return inst.releasePool
}

func (inst *contextRuntime) GetArguments() collection.Arguments {
	args := inst.arguments
	return args
}

func (inst *contextRuntime) GetAttributes() collection.Attributes {
	return inst.attributes
}

func (inst *contextRuntime) GetEnvironment() collection.Environment {
	return inst.environment
}

func (inst *contextRuntime) GetProperties() collection.Properties {
	return inst.properties
}

func (inst *contextRuntime) GetParameters() collection.Parameters {
	return inst.parameters
}

func (inst *contextRuntime) GetResources() collection.Resources {
	return inst.resources
}

func (inst *contextRuntime) GetApplicationName() string {
	return inst.appName
}

func (inst *contextRuntime) GetApplicationVersion() string {
	return inst.appVersion
}

func (inst *contextRuntime) GetStartupTimestamp() int64 {
	return inst.time1
}

func (inst *contextRuntime) GetShutdownTimestamp() int64 {
	return inst.time2
}

func (inst *contextRuntime) GetURI() string {
	return inst.uri
}

func (inst *contextRuntime) GetErrorHandler() lang.ErrorHandler {
	h := inst.errorHandler
	if h == nil {
		h = lang.DefaultErrorHandler()
	}
	return h
}

func (inst *contextRuntime) SetErrorHandler(h lang.ErrorHandler) {
	inst.errorHandler = h
}

func (inst *contextRuntime) FindComponent(selector string) (lang.Object, error) {
	holder, err := inst.GetComponents().GetComponent(selector)
	if err != nil {
		return nil, err
	}
	if inst.isComponentReady(holder) {
		o := holder.GetInstance().Get()
		return o, nil
	}
	scope := application.ScopePrototype
	ccc := inst.openCreationContext(scope)
	o, err := ccc.LoadComponent(holder)
	err = ccc.Close()
	return o, err
}

func (inst *contextRuntime) FindComponents(selector string) []lang.Object {
	holders := inst.GetComponents().GetComponents(selector)
	if inst.isAllComponentsReady(holders) {
		results := make([]lang.Object, 0)
		for index := range holders {
			h := holders[index]
			if h == nil {
				continue
			}
			o := h.GetInstance().Get()
			results = append(results, o)
		}
		return results
	}
	scope := application.ScopePrototype
	ccc := inst.openCreationContext(scope)
	results := ccc.LoadComponents(holders)
	ccc.Close()
	return results
}

func (inst *contextRuntime) Injector() application.Injector {
	return inst.InjectorScope(application.ScopePrototype)
}

func (inst *contextRuntime) InjectorScope(scope application.ComponentScope) application.Injector {
	ccc := inst.openCreationContext(scope)
	injector := &innerInjector{}
	injector.init(ccc, true)
	return injector
}

func (inst *contextRuntime) openCreationContext(scope application.ComponentScope) CreationContext {
	creation := &contextCreation{}
	return creation.init(inst, scope)
}

func (inst *contextRuntime) isComponentReady(h application.ComponentHolder) bool {
	if h == nil {
		return false
	}
	info := h.GetInfo()
	scope := info.GetScope()
	if scope == application.ScopeSingleton {
		instance := h.GetInstance()
		return instance.IsLoaded()
	} else if scope == application.ScopePrototype {
		return false
	} else {
		return false
	}
}

func (inst *contextRuntime) isAllComponentsReady(list []application.ComponentHolder) bool {
	if list == nil {
		return true
	}
	for index := range list {
		item := list[index]
		if item == nil {
			continue
		}
		if !inst.isComponentReady(item) {
			return false
		} else {
			continue
		}
	}
	return true
}

func (inst *contextRuntime) NewChild() application.Context {
	child := &contextRuntime{}
	_, err := child.Init(inst)
	if err == nil {
		return child
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
