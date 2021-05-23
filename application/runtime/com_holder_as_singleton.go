package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

type SingletonComponentHolder struct {
	context   application.RuntimeContext
	info      application.ComponentInfo
	singleton application.ComponentInstance
}

////////////////////////////////////////////////////////////////////////////////
// impl SingletonComponentHolder

func (inst *SingletonComponentHolder) GetInstance() application.ComponentInstance {
	single := inst.singleton
	if single == nil {
		single = inst.info.GetFactory().NewInstance()
		inst.singleton = single
	}
	return single
}

func (inst *SingletonComponentHolder) IsOriginalName(name string) bool {
	return (name == inst.info.GetID())
}

func (inst *SingletonComponentHolder) GetInfo() application.ComponentInfo {
	return inst.info
}

func (inst *SingletonComponentHolder) GetContext() application.RuntimeContext {
	return inst.context
}

func (inst *SingletonComponentHolder) GetPrototype() lang.Object {
	return inst.info.GetPrototype()
}

func (inst *SingletonComponentHolder) MakeChild(context application.RuntimeContext) application.ComponentHolder {

	singleton := inst.GetInstance()

	if context == nil {
		context = inst.context
	}

	child := &SingletonComponentHolder{}
	child.info = inst.info
	child.context = context
	child.singleton = singleton
	return child
}
