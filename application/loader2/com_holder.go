package loader2

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////

// 单例 holder
type singletonComponentHolder struct {
	info      application.ComponentInfo
	context   application.Context
	singleton application.ComponentInstance
}

func (inst *singletonComponentHolder) _Impl() application.ComponentHolder {
	return inst
}

func (inst *singletonComponentHolder) init(ctx application.Context, i application.ComponentInstance, info application.ComponentInfo) application.ComponentHolder {
	if i == nil {
		i = info.GetFactory().NewInstance()
	}
	inst.context = ctx
	inst.singleton = i
	inst.info = info
	return inst
}

func (inst *singletonComponentHolder) GetInstance() application.ComponentInstance {
	return inst.singleton
}

func (inst *singletonComponentHolder) IsOriginalName(name string) bool {
	return name == inst.info.GetID()
}

func (inst *singletonComponentHolder) GetInfo() application.ComponentInfo {
	return inst.info
}

func (inst *singletonComponentHolder) GetPrototype() lang.Object {
	return inst.info.GetPrototype()
}

func (inst *singletonComponentHolder) GetContext() application.Context {
	return inst.context
}

func (inst *singletonComponentHolder) MakeChild(context application.Context) application.ComponentHolder {
	if context == nil {
		context = inst.context
	}
	child := &singletonComponentHolder{}
	return child.init(context, inst.singleton, inst.info)
}

////////////////////////////////////////////////////////////////////////////////

// 原型 holder
type prototypeComponentHolder struct {
	info    application.ComponentInfo
	context application.Context
}

func (inst *prototypeComponentHolder) _Impl() application.ComponentHolder {
	return inst
}

func (inst *prototypeComponentHolder) init(ctx application.Context, info application.ComponentInfo) application.ComponentHolder {
	inst.context = ctx
	inst.info = info
	return inst
}

func (inst *prototypeComponentHolder) GetInstance() application.ComponentInstance {
	return inst.info.GetFactory().NewInstance()
}

func (inst *prototypeComponentHolder) IsOriginalName(name string) bool {
	return name == inst.info.GetID()
}

func (inst *prototypeComponentHolder) GetInfo() application.ComponentInfo {
	return inst.info
}

func (inst *prototypeComponentHolder) GetPrototype() lang.Object {
	return inst.info.GetPrototype()
}

func (inst *prototypeComponentHolder) GetContext() application.Context {
	return inst.context
}

func (inst *prototypeComponentHolder) MakeChild(context application.Context) application.ComponentHolder {
	if context == nil {
		context = inst.context
	}
	child := &prototypeComponentHolder{}
	return child.init(context, inst.info)
}

////////////////////////////////////////////////////////////////////////////////
