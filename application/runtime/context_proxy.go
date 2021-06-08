package runtime

import (
	"io"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

type ContextProxy struct {
	io.Closer

	Current  application.Context
	Creation application.Context
	Runtime  application.Context
}

func (inst *ContextProxy) Close() error {
	inst.Current = inst.Runtime
	inst.Creation = nil
	return nil
}

func (inst *ContextProxy) GetComponent(selector string) (lang.Object, error) {
	return inst.Current.GetComponent(selector)
}

func (inst *ContextProxy) GetComponentList(selector string) ([]lang.Object, error) {
	return inst.Current.GetComponentList(selector)
}

func (inst *ContextProxy) Injector() application.Injector {
	return inst.Current.Injector()
}

func (inst *ContextProxy) ComponentLoader() application.ComponentLoader {
	return inst.Current.ComponentLoader()
}

func (inst *ContextProxy) GetComponents() application.Components {
	return inst.Current.GetComponents()
}

func (inst *ContextProxy) GetReleasePool() lang.ReleasePool {
	return inst.Current.GetReleasePool()
}

func (inst *ContextProxy) GetArguments() collection.Arguments {
	return inst.Current.GetArguments()
}

func (inst *ContextProxy) GetAttributes() collection.Attributes {
	return inst.Current.GetAttributes()
}

func (inst *ContextProxy) GetEnvironment() collection.Environment {
	return inst.Current.GetEnvironment()
}

func (inst *ContextProxy) GetProperties() collection.Properties {
	return inst.Current.GetProperties()
}

func (inst *ContextProxy) GetParameters() collection.Parameters {
	return inst.Current.GetParameters()
}

func (inst *ContextProxy) GetResources() collection.Resources {
	return inst.Current.GetResources()
}

func (inst *ContextProxy) GetApplicationName() string {
	return inst.Current.GetApplicationName()
}

func (inst *ContextProxy) GetURI() string {
	return inst.Current.GetURI()
}

func (inst *ContextProxy) GetApplicationVersion() string {
	return inst.Current.GetApplicationVersion()
}

func (inst *ContextProxy) GetStartupTimestamp() int64 {
	return inst.Current.GetStartupTimestamp()
}

func (inst *ContextProxy) GetShutdownTimestamp() int64 {
	return inst.Current.GetShutdownTimestamp()
}

func (inst *ContextProxy) NewChild() application.Context {
	return inst.Current.NewChild()
}

func (inst *ContextProxy) GetErrorHandler() lang.ErrorHandler {
	return inst.Current.GetErrorHandler()
}

func (inst *ContextProxy) SetErrorHandler(h lang.ErrorHandler) {
	inst.Current.SetErrorHandler(h)
}
