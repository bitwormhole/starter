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

// comInfo ????????????????????? ComponentInfo ??????
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

// GetAliases ?????????????????????
func (inst *comInfo) GetAliases() []string {
	return inst.stringToItems(inst.Aliases)
}

// GetClasses ?????????????????????
func (inst *comInfo) GetClasses() []string {
	return inst.stringToItems(inst.Class)
}

// GetID ???????????????ID
func (inst *comInfo) GetID() string {
	return inst.ID
}

// GetScope ????????????????????????
func (inst *comInfo) GetScope() application.ComponentScope {
	return inst.Scope
}

// GetFactory ?????????????????????
func (inst *comInfo) GetFactory() application.ComponentFactory {
	return inst.factory
}

// GetPrototype ?????????????????????
func (inst *comInfo) GetPrototype() lang.Object {
	return inst.factory.GetPrototype()
}

// IsTypeOf ?????? class ????????????
func (inst *comInfo) IsTypeOf(name string) bool {
	// nop
	panic("no impl")
	return false
}

// IsNameOf ?????? id(or alias) ????????????
func (inst *comInfo) IsNameOf(name string) bool {
	// nop
	panic("no impl")
	return false
}
