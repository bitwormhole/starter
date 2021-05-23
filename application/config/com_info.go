package config

import (
	"errors"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

// ComInfo 提供一个简易的 ComponentInfo 实现
type ComInfo struct {
	// implements ComponentInfo
	ID      string
	Class   string
	Scope   application.ComponentScope
	Aliases []string

	OnNew     func() lang.Object
	OnInit    func(obj lang.Object) error
	OnDestroy func(obj lang.Object) error
	OnInject  func(obj lang.Object, context application.Context) error

	demo application.ComponentInfo
}

type comFactoryForComInfo struct {
	info *ComInfo
}

type comInstanceForComInfo struct {
	info   *ComInfo
	target lang.Object
}

////////////////////////////////////////////////////////////////////////////////

func (inst *comFactoryForComInfo) NewInstance() application.ComponentInstance {
	instance := &comInstanceForComInfo{}
	err := instance.initial(inst.info)
	if err != nil {
		panic(err)
	}
	return instance
}

////////////////////////////////////////////////////////////////////////////////

func (inst *comInstanceForComInfo) initial(info *ComInfo) error {
	fnNew := info.OnNew
	if fnNew == nil {
		return errors.New("no func:OnNew for component:" + info.GetID())
	}
	inst.info = info
	inst.target = fnNew()
	return nil
}

func (inst *comInstanceForComInfo) Get() lang.Object {
	return inst.target
}

func (inst *comInstanceForComInfo) IsLoaded() bool {
	return false
}

func (inst *comInstanceForComInfo) Inject(context application.Context) error {
	fnInject := inst.info.OnInject
	if fnInject == nil {
		return nil
	}
	return fnInject(inst.target, context)
}

func (inst *comInstanceForComInfo) Init() error {
	fn := inst.info.OnInit
	if fn == nil {
		return nil
	}
	return fn(inst.target)
}

func (inst *comInstanceForComInfo) Destroy() error {
	fn := inst.info.OnDestroy
	if fn == nil {
		return nil
	}
	return fn(inst.target)
}

////////////////////////////////////////////////////////////////////////////////

// GetAliases 获取组件的别名
func (inst *ComInfo) GetAliases() []string {
	return inst.Aliases
}

// GetAliases 获取组件的别名
func (inst *ComInfo) GetClasses() []string {
	return inst.Aliases
}

// GetID 获取组件的ID
func (inst *ComInfo) GetID() string {
	return inst.ID
}

// GetClass 获取组件的类
func (inst *ComInfo) GetClass() string {
	return inst.Class
}

// GetScope 获取组件的作用域
func (inst *ComInfo) GetScope() application.ComponentScope {
	return inst.Scope
}

// GetFactory 获取组件的工厂
func (inst *ComInfo) GetFactory() application.ComponentFactory {
	return &comFactoryForComInfo{info: inst}
}

// GetPrototype 获取组件的原型
func (inst *ComInfo) GetPrototype() lang.Object {
	return nil
}

// IsTypeOf 判断 class 是否匹配
func (inst *ComInfo) IsTypeOf(name string) bool {
	// nop
	return false
}

// IsNameOf 判断 id(or alias) 是否匹配
func (inst *ComInfo) IsNameOf(name string) bool {
	// nop
	return false
}
