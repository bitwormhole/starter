package config

import (
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// config:com-factory

// type comFactoryForComInfo struct {
// 	info *comInfo
// }

// func (inst *comFactoryForComInfo) _Impl() (application.ComponentFactory, application.ComponentAfterService) {
// 	return inst, inst
// }

// func (inst *comFactoryForComInfo) GetPrototype() lang.Object {
// 	pt := inst.info.prototype
// 	if pt == nil {
// 		pt = inst.NewInstance().Get()
// 		inst.info.prototype = pt
// 	}
// 	return pt
// }

// func (inst *comFactoryForComInfo) AfterService() application.ComponentAfterService {
// 	return inst
// }

// func (inst *comFactoryForComInfo) Inject(i application.ComponentInstance, ctx application.InstanceContext) error {
// 	return inst.info.innerFactory.AfterService().Inject(i, ctx)
// }

// func (inst *comFactoryForComInfo) Init(i application.ComponentInstance) error {
// 	return inst.info.innerFactory.AfterService().Init(i)
// }

// func (inst *comFactoryForComInfo) Destroy(i application.ComponentInstance) error {
// 	return inst.info.innerFactory.AfterService().Destroy(i)
// }

// func (inst *comFactoryForComInfo) NewInstance() application.ComponentInstance {
// 	instance := &comInstanceForComInfo{}
// 	err := instance.initial(inst.info)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return instance
// }

////////////////////////////////////////////////////////////////////////////////
// config:com-instance

// type comInstanceForComInfo struct {
// 	info   *comInfo
// 	target lang.Object
// }

// func (inst *comInstanceForComInfo) _Impl() application.ComponentInstance {
// 	return inst
// }

// func (inst *comInstanceForComInfo) State() application.ComponentState {
// 	return application.StateZero
// }

// func (inst *comInstanceForComInfo) Factory() application.ComponentFactory {
// 	return inst.info.factory
// }

// func (inst *comInstanceForComInfo) initial(info *comInfo) error {
// 	inst.info = info
// 	inst.target = info.factory.NewInstance().Get()
// 	return nil
// }

// func (inst *comInstanceForComInfo) Get() lang.Object {
// 	return inst.target
// }

// func (inst *comInstanceForComInfo) IsLoaded() bool {
// 	return false
// }

// func (inst *comInstanceForComInfo) Inject(ictx application.InstanceContext) error {
// 	return inst.info.innerFactory.AfterService().Inject(inst, ictx)
// }

// func (inst *comInstanceForComInfo) Init() error {
// 	return inst.info.innerFactory.AfterService().Init(inst)
// }

// func (inst *comInstanceForComInfo) Destroy() error {
// 	return inst.info.innerFactory.AfterService().Destroy(inst)
// }

////////////////////////////////////////////////////////////////////////////////
// config:com-info

// comInfo 提供一个简易的 ComponentInfo 实现
type comInfo struct {
	// implements ComponentInfo
	ID      string
	Class   string
	Aliases string
	Scope   application.ComponentScope

	// OnNew     OnNew
	// OnInject  OnInject
	// OnInit    OnInit
	// OnDestroy OnDestroy

	// prototype lang.Object
	// innerFactory application.ComponentFactory
	factory application.ComponentFactory
}

func (inst *comInfo) init() {

	// //	factory := &comFactoryForComInfo{}
	// 	factory.info = inst
	// 	inst.factory = factory
	// 	inst.prototype = factory.NewInstance().Get()
	// return
}

func (inst *comInfo) stringToItems(strlist string) []string {
	strlist = strings.ReplaceAll(strlist, "\t", ",")
	strlist = strings.ReplaceAll(strlist, " ", ",")
	array := strings.Split(strlist, ",")
	results := make([]string, 0)
	for _, item := range array {
		if item == "" {
			continue
		}
		results = append(results, item)
	}
	return results
}

// GetAliases 获取组件的别名
func (inst *comInfo) GetAliases() []string {
	return inst.stringToItems(inst.Aliases)
}

// GetClasses 获取组件的别名
func (inst *comInfo) GetClasses() []string {
	return inst.stringToItems(inst.Class)
}

// GetID 获取组件的ID
func (inst *comInfo) GetID() string {
	return inst.ID
}

// GetScope 获取组件的作用域
func (inst *comInfo) GetScope() application.ComponentScope {
	return inst.Scope
}

// GetFactory 获取组件的工厂
func (inst *comInfo) GetFactory() application.ComponentFactory {
	return inst.factory
}

// GetPrototype 获取组件的原型
func (inst *comInfo) GetPrototype() lang.Object {
	return inst.factory.GetPrototype()
}

// IsTypeOf 判断 class 是否匹配
func (inst *comInfo) IsTypeOf(name string) bool {
	// nop
	panic("no impl")
	return false
}

// IsNameOf 判断 id(or alias) 是否匹配
func (inst *comInfo) IsNameOf(name string) bool {
	// nop
	panic("no impl")
	return false
}
