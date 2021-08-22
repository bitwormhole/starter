package config

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

func SimpleInstance(f application.ComponentFactory, i lang.Object) application.ComponentInstance {
	ci := &simpleInstance{}
	ci.factory = f
	ci.target = i
	return ci
}

////////////////////////////////////////////////////////////////////////////////

type simpleInstance struct {
	factory application.ComponentFactory
	target  lang.Object
}

func (inst *simpleInstance) _Impl() application.ComponentInstance {
	return inst
}

func (inst *simpleInstance) Factory() application.ComponentFactory {
	return inst.factory
}

func (inst *simpleInstance) Get() lang.Object {
	return inst.target
}

func (inst *simpleInstance) State() application.ComponentState {
	panic("unsupported")
}

func (inst *simpleInstance) Inject(context application.InstanceContext) error {
	panic("unsupported")
}

func (inst *simpleInstance) Init() error {
	panic("unsupported")
}

func (inst *simpleInstance) Destroy() error {
	panic("unsupported")
}

////////////////////////////////////////////////////////////////////////////////
