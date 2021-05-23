package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

type contextProxy struct {
	current  application.Context
	creation application.Context
	runtime  application.Context
	pool     lang.ReleasePool
}

func (inst *contextProxy) start() application.Context {
	inst.current = inst.runtime
	inst.creation = nil
	return inst
}

func (inst *contextProxy) FindComponent(selector string) (lang.Object, error) {
	return inst.current.FindComponent(selector)
}

func (inst *contextProxy) FindComponents(selector string) []lang.Object {
	return inst.current.FindComponents(selector)
}

func (inst *contextProxy) Injector() application.Injector {
	return inst.current.Injector()
}

func (inst *contextProxy) InjectorScope(scope application.ComponentScope) application.Injector {
	return inst.current.InjectorScope(scope)
}

func (inst *contextProxy) GetComponents() application.Components {
	return inst.current.GetComponents()
}

func (inst *contextProxy) GetReleasePool() lang.ReleasePool {
	pool := inst.pool
	if pool == nil {
		pool = inst.current.GetReleasePool()
	}
	return pool
}

func (inst *contextProxy) GetArguments() collection.Arguments {
	return inst.current.GetArguments()
}

func (inst *contextProxy) GetAttributes() collection.Attributes {
	return inst.current.GetAttributes()
}

func (inst *contextProxy) GetEnvironment() collection.Environment {
	return inst.current.GetEnvironment()
}

func (inst *contextProxy) GetProperties() collection.Properties {
	return inst.current.GetProperties()
}

func (inst *contextProxy) GetParameters() collection.Parameters {
	return inst.current.GetParameters()
}

func (inst *contextProxy) GetResources() collection.Resources {
	return inst.current.GetResources()
}

func (inst *contextProxy) GetApplicationName() string {
	return inst.current.GetApplicationName()
}

func (inst *contextProxy) GetURI() string {
	return inst.current.GetURI()
}

func (inst *contextProxy) GetApplicationVersion() string {
	return inst.current.GetApplicationVersion()
}

func (inst *contextProxy) GetStartupTimestamp() int64 {
	return inst.current.GetStartupTimestamp()
}

func (inst *contextProxy) GetShutdownTimestamp() int64 {
	return inst.current.GetShutdownTimestamp()
}

func (inst *contextProxy) NewChild() application.Context {
	return inst.current.NewChild()
}

func (inst *contextProxy) GetErrorHandler() lang.ErrorHandler {
	return inst.current.GetErrorHandler()
}

func (inst *contextProxy) SetErrorHandler(h lang.ErrorHandler) {
	inst.current.SetErrorHandler(h)
}
