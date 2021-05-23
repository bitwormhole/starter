package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct

type componentLoading struct {
	holder       application.ComponentHolder
	instance     application.ComponentInstance
	loadingOrder int
	hasStarted   bool
	hasInjected  bool
}

type componentLoadingSorter struct {
	items []*componentLoading
}

////////////////////////////////////////////////////////////////////////////////
// struct componentInstanceCloser

type componentInstanceCloser struct {
	instance application.ComponentInstance
}

func (inst *componentInstanceCloser) Dispose() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// impl creationComponentLoader

func (inst *componentLoading) tryStart(pool lang.ReleasePool) error {
	if inst.hasStarted {
		return nil
	} else {
		inst.hasStarted = true
	}
	if inst.instance.IsLoaded() {
		return nil
	}
	err := inst.instance.Init()
	if err != nil {
		return err
	}
	pool.Push(&componentInstanceCloser{instance: inst.instance})
	return nil
}

func (inst *componentLoading) tryInject(context application.Context) error {
	if inst.hasInjected {
		return nil
	} else {
		inst.hasInjected = true
	}
	if inst.instance.IsLoaded() {
		return nil
	}
	return inst.instance.Inject(context)
}

////////////////////////////////////////////////////////////////////////////////
// impl componentLoadingSorter

func (inst *componentLoadingSorter) Len() int {
	return len(inst.items)
}

func (inst *componentLoadingSorter) Less(i, j int) bool {
	a := inst.items[i]
	b := inst.items[j]
	return a.loadingOrder > b.loadingOrder
}

func (inst *componentLoadingSorter) Swap(i, j int) {
	a := inst.items[i]
	b := inst.items[j]
	inst.items[i] = b
	inst.items[j] = a
}

////////////////////////////////////////////////////////////////////////////////
// EOF
