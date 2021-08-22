package loader2

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

type comInstanceWrapper struct {
	factory application.ComponentFactory
	target  lang.Object
	state   application.ComponentState
	pool    lang.ReleasePool
}

func (inst *comInstanceWrapper) init(i application.ComponentInstance, f application.ComponentFactory) application.ComponentInstance {
	inst.target = i.Get()
	inst.factory = f
	inst.state = application.StateZero
	inst.pool = nil
	return inst
}

func (inst *comInstanceWrapper) Factory() application.ComponentFactory {
	return inst.factory
}

func (inst *comInstanceWrapper) Get() lang.Object {
	return inst.target
}

func (inst *comInstanceWrapper) State() application.ComponentState {
	return inst.state
}

func (inst *comInstanceWrapper) Inject(context application.InstanceContext) error {
	const want = application.StateInjected
	if inst.state >= want {
		return nil
	}
	inst.pool = context.Pool()
	inst.state = want
	return inst.factory.AfterService().Inject(inst, context)
}

func (inst *comInstanceWrapper) Init() error {
	const want = application.StateInitialled
	if inst.state >= want {
		return nil
	}
	inst.state = want
	err := inst.factory.AfterService().Init(inst)
	if err != nil {
		return err
	}
	inst.pool.Push(inst.toDisposable())
	return nil
}

func (inst *comInstanceWrapper) Destroy() error {
	const want = application.StateDestroyed
	if inst.state >= want {
		return nil
	}
	inst.state = want
	return inst.factory.AfterService().Destroy(inst)
}

func (inst *comInstanceWrapper) toDisposable() lang.Disposable {
	return inst
}

func (inst *comInstanceWrapper) Dispose() error {
	return inst.Destroy()
}
