package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

type contextProxy struct {
	current  application.RuntimeContext
	creation application.RuntimeContext
	runtime  application.RuntimeContext
}

func (inst *contextProxy) startRuntime() {
	inst.current = inst.runtime
	inst.creation = nil
}

func (inst *contextProxy) GetComponents() application.Components {
	return inst.current.GetComponents()
}

func (inst *contextProxy) GetReleasePool() collection.ReleasePool {
	return inst.current.GetReleasePool()
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

func (inst *contextProxy) NewChild() application.RuntimeContext {
	return inst.current.NewChild()
}

func (inst *contextProxy) GetErrorHandler() lang.ErrorHandler {
	return inst.current.GetErrorHandler()
}

func (inst *contextProxy) SetErrorHandler(h lang.ErrorHandler) {
	inst.current.SetErrorHandler(h)
}

func (inst *contextProxy) OpenCreationContext(scope application.ComponentScope) application.CreationContext {
	return inst.current.OpenCreationContext(scope)
}
