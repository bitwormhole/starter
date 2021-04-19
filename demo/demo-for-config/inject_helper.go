package config

import "github.com/bitwormhole/starter/application"

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

func (inst *injectHelper) getCar(name string) *Car {
	obj, err := inst.context.GetComponents().GetComponent(name)
	if err == nil {
		return obj.(*Car)
	}
	inst.handleError(err)
	return nil
}

func (inst *injectHelper) getEngine(name string) *Engine {
	obj, err := inst.context.GetComponents().GetComponent(name)
	if err == nil {
		return obj.(*Engine)
	}
	inst.handleError(err)
	return nil
}
