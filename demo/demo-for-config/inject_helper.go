package config

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/demo/demo-for-config/components"
)

type injectHelper struct {
	context application.RuntimeContext
	err     error
}

func (inst *injectHelper) init(ctx application.RuntimeContext) {
	inst.context = ctx
}

func (inst *injectHelper) handleError(err error) {
	if err == nil {
		return
	}
	inst.err = err
}

func (inst *injectHelper) getCar(name string) *components.Car {
	obj, err := inst.context.GetComponents().GetComponent(name)
	if err == nil {
		return obj.(*components.Car)
	}
	inst.handleError(err)
	return nil
}

func (inst *injectHelper) getEngine(name string) *components.Engine {
	obj, err := inst.context.GetComponents().GetComponent(name)
	if err == nil {
		return obj.(*components.Engine)
	}
	inst.handleError(err)
	return nil
}
