package loader2

import (
	"fmt"
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/util"
)

type appContextInner struct {
	deadline time.Time
	err      error
}

type appContext struct {
	inner appContextInner

	appName           string
	appTitle          string
	appURI            string
	appVersion        string
	startupTimestamp  int64
	shutdownTimestamp int64

	errorHandler lang.ErrorHandler
	pool         lang.ReleasePool
	components   application.Components
	closed       bool

	arguments   collection.Arguments
	attributes  collection.Attributes
	environment collection.Environment
	parameters  collection.Parameters
	properties  collection.Properties
	resources   collection.Resources
}

func (inst *appContext) init() application.Context {

	comMan := &componentManager{}
	comMan.init(inst)

	inst.errorHandler = &defaultErrorHandler{}
	inst.pool = lang.CreateReleasePool()
	inst.components = comMan

	inst.arguments = collection.CreateArguments()
	inst.attributes = collection.CreateAttributes()
	inst.environment = collection.CreateEnvironment()
	inst.parameters = collection.CreateParameters()
	inst.properties = collection.CreateProperties()
	inst.resources = collection.CreateResources()

	inst.appName = "Starter"
	inst.appTitle = "Starter Application"
	inst.appVersion = "v0.0.0"
	inst.appURI = "https://bitwormhole.com/starter"
	inst.startupTimestamp = util.CurrentTimestamp()
	inst.shutdownTimestamp = 0

	inst.pool.Push(lang.DisposableForFunc(func() error { return inst.onShutdown() }))

	return inst
}

func (inst *appContext) onShutdown() error {
	inst.shutdownTimestamp = util.CurrentTimestamp()
	return nil
}

func (inst *appContext) GetURI() string {
	return inst.appURI
}

func (inst *appContext) GetApplicationName() string {
	return inst.appName
}

func (inst *appContext) GetApplicationVersion() string {
	return inst.appVersion
}

func (inst *appContext) GetStartupTimestamp() int64 {
	return inst.startupTimestamp
}

func (inst *appContext) GetShutdownTimestamp() int64 {
	return inst.shutdownTimestamp
}

// GetReleasePool 取context的生命周期管理池
func (inst *appContext) GetReleasePool() lang.ReleasePool {
	return inst.pool
}

// GetComponents 取context组件管理器
func (inst *appContext) GetComponents() application.Components {
	return inst.components
}

// GetResources 取context的资源管理器
func (inst *appContext) GetResources() collection.Resources {
	return inst.resources
}

func (inst *appContext) GetArguments() collection.Arguments {
	return inst.arguments
}

func (inst *appContext) GetAttributes() collection.Attributes {
	return inst.attributes
}

func (inst *appContext) GetEnvironment() collection.Environment {
	return inst.environment
}

func (inst *appContext) GetProperties() collection.Properties {
	return inst.properties
}

func (inst *appContext) GetParameters() collection.Parameters {
	return inst.parameters
}

func (inst *appContext) SetErrorHandler(h lang.ErrorHandler) {
	if h == nil {
		return
	}
	inst.errorHandler = h
}

func (inst *appContext) GetErrorHandler() lang.ErrorHandler {
	return inst.errorHandler
}

func (inst *appContext) NewChild() application.Context {

	child := &appContext{}
	child.init()

	child.errorHandler = inst.errorHandler

	child.arguments.Import(inst.arguments.Export())
	child.attributes.Import(inst.attributes.Export(nil))
	child.environment.Import(inst.environment.Export(nil))
	child.parameters.Import(inst.parameters.Export(nil))
	child.properties.Import(inst.properties.Export(nil))
	child.resources.Import(inst.resources.Export(nil), false)

	child.components.Import(inst.components.Export(nil))

	child.appName = inst.appName
	child.appTitle = inst.appTitle
	child.appURI = inst.appURI
	child.appVersion = inst.appVersion

	child.components.GroupManager().Reload()

	return child
}

func (inst *appContext) Close() error {
	if inst.closed {
		return nil
	}
	inst.closed = true
	return inst.pool.Release()
}

func (inst *appContext) isComReady(ci application.ComponentInstance) bool {
	if ci == nil {
		return false
	}
	state := ci.State()
	return application.StateInitialled <= state && state < application.StateDestroyed
}

func (inst *appContext) isAllComReady(list []application.ComponentInstance) bool {
	if list == nil {
		return true
	}
	for _, item := range list {
		if inst.isComReady(item) {
			continue
		}
		return false
	}
	return true
}

func (inst *appContext) GetComponent(selector string) (lang.Object, error) {

	h, err := inst.GetComponents().FindComponent(selector)
	if err != nil {
		return nil, err
	}

	i := h.GetInstance()
	if inst.isComReady(i) {
		return i.Get(), nil
	}

	loader := &comInstanceLoader{}
	loader.init(inst, true)
	loader.addInstance(i)
	err = loader.Load()
	if err != nil {
		return nil, err
	}

	return i.Get(), nil
}

func (inst *appContext) GetComponentList(selector string) ([]lang.Object, error) {

	holders := inst.GetComponents().FindComponents(selector)
	inslist := make([]application.ComponentInstance, 0)
	objlist := make([]lang.Object, 0)

	for _, h := range holders {
		i := h.GetInstance()
		inslist = append(inslist, i)
	}

	for _, i := range inslist {
		objlist = append(objlist, i.Get())
	}

	if inst.isAllComReady(inslist) {
		return objlist, nil
	}

	loader := &comInstanceLoader{}
	loader.init(inst, true)
	loader.addInstances(inslist)
	err := loader.Load()
	if err != nil {
		return nil, err
	}

	return objlist, nil
}

func (inst *appContext) Err() error {
	return inst.inner.err
}

func (inst *appContext) Done() <-chan struct{} {
	return nil
}

func (inst *appContext) Deadline() (time.Time, bool) {
	return inst.inner.deadline, false
}

func (inst *appContext) Value(key interface{}) interface{} {
	name, ok := key.(string)
	if !ok {
		name = fmt.Sprint(key)
	}
	return inst.GetValue(name)
}

func (inst *appContext) GetValue(key string) interface{} {
	return inst.GetAttributes().GetAttribute(key)
}

func (inst *appContext) SetValue(key string, value interface{}) {
	inst.GetAttributes().SetAttribute(key, value)
}
