package runtime

import (
	"github.com/bitwormhole/starter/application"
)

type PrototypeComponentHolder struct {
	context application.RuntimeContext
	info    application.ComponentInfo
}

////////////////////////////////////////////////////////////////////////////////
// impl PrototypeComponentHolder

func (inst *PrototypeComponentHolder) IsOriginalName(name string) bool {
	return (name == inst.info.GetID())
}

func (inst *PrototypeComponentHolder) GetInstance() application.ComponentInstance {
	factory := inst.info.GetFactory()
	return factory.NewInstance()
}

func (inst *PrototypeComponentHolder) GetInfo() application.ComponentInfo {
	return inst.info
}

func (inst *PrototypeComponentHolder) GetContext() application.RuntimeContext {
	return inst.context
}

func (inst *PrototypeComponentHolder) MakeChild(context application.RuntimeContext) application.ComponentHolder {

	if context == nil {
		context = inst.context
	}

	child := &PrototypeComponentHolder{}
	child.context = context
	child.info = inst.info
	return child
}
