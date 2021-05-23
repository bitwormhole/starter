package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

type comInstanceWithState struct {
	// data
	inner   application.ComponentInstance
	context application.Context
	target  lang.Object

	// state
	injectErr  error
	initErr    error
	destroyErr error

	injectDone  bool
	initDone    bool
	destroyDone bool
}

func createComInstanceWithState(inner application.ComponentInstance) application.ComponentInstance {
	inst := &comInstanceWithState{}
	inst.target = inner.Get()
	inst.inner = inner
	return inst
}

func (inst *comInstanceWithState) IsLoaded() bool {
	return (inst.injectDone && inst.initDone)
}

func (inst *comInstanceWithState) Get() lang.Object {
	return inst.target
}

func (inst *comInstanceWithState) Inject(context application.Context) error {
	if inst.injectDone {
		return inst.injectErr
	}
	err := inst.inner.Inject(context)
	inst.injectDone = true
	inst.injectErr = err
	inst.context = context
	return err
}

func (inst *comInstanceWithState) Init() error {
	if inst.initDone {
		return inst.initErr
	}
	err := inst.inner.Init()
	inst.initDone = true
	inst.initErr = err
	if err == nil {
		pool := inst.context.GetReleasePool()
		pool.Push(&comInstanceWithStateCloser{target: inst})
	}
	return err
}

func (inst *comInstanceWithState) Destroy() error {
	if inst.destroyDone {
		return inst.destroyErr
	}
	err := inst.inner.Destroy()
	inst.destroyDone = true
	inst.destroyErr = err
	return err
}

////////////////////////////////////////////////////////////////////////////////
// comInstanceWithStateFactory class
type comInstanceWithStateFactory struct {
	inner application.ComponentFactory
}

func createComInstanceWithStateFactory(f application.ComponentFactory) application.ComponentFactory {
	inst := &comInstanceWithStateFactory{inner: f}
	return inst
}

func (inst *comInstanceWithStateFactory) NewInstance() application.ComponentInstance {
	innerInstance := inst.inner.NewInstance()
	return createComInstanceWithState(innerInstance)
}

////////////////////////////////////////////////////////////////////////////////

type comInstanceWithStateCloser struct {
	target *comInstanceWithState
}

func (inst *comInstanceWithStateCloser) Dispose() error {
	return inst.target.Destroy()
}
