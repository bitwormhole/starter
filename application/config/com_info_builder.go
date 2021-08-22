package config

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter/application"
)

type comInfoBuilder struct {
	id      string
	class   string
	aliases string
	scope   string

	// fnNew     OnNew
	// fnInject  OnInject
	// fnInit    OnInit
	// fnDestroy OnDestroy

	simpleFactory application.ComponentFactory
}

func (inst *comInfoBuilder) _Impl() ComponentInfoBuilder {
	return inst
}

func (inst *comInfoBuilder) Next() ComponentInfoBuilder {
	inst.Reset()
	return inst
}

func (inst *comInfoBuilder) ID(id string) ComponentInfoBuilder {
	inst.id = id
	return inst
}

func (inst *comInfoBuilder) Class(cls string) ComponentInfoBuilder {
	inst.class = cls
	return inst
}

func (inst *comInfoBuilder) Aliases(aliases string) ComponentInfoBuilder {
	inst.aliases = aliases
	return inst
}

func (inst *comInfoBuilder) Scope(scope string) ComponentInfoBuilder {
	inst.scope = scope
	return inst
}

// func (inst *comInfoBuilder) OnNew(fn OnNew) ComponentInfoBuilder {
// 	inst.fnNew = fn
// 	return inst
// }

// func (inst *comInfoBuilder) OnInject(fn OnInject) ComponentInfoBuilder {
// 	inst.fnInject = fn
// 	return inst
// }

// func (inst *comInfoBuilder) OnInit(fn OnInit) ComponentInfoBuilder {
// 	inst.fnInit = fn
// 	return inst
// }

// func (inst *comInfoBuilder) OnDestroy(fn OnDestroy) ComponentInfoBuilder {
// 	inst.fnDestroy = fn
// 	return inst
// }

func (inst *comInfoBuilder) Factory(f application.ComponentFactory) ComponentInfoBuilder {
	inst.simpleFactory = f
	return inst
}

func (inst *comInfoBuilder) Reset() {

	inst.id = ""
	inst.class = ""
	inst.scope = ""
	inst.aliases = ""
	inst.simpleFactory = nil
}

func (inst *comInfoBuilder) Create() (application.ComponentInfo, error) {

	scope, err := inst.parseScope(inst.scope)
	if err != nil {
		return nil, err
	}

	info := &comInfo{}
	info.ID = inst.id
	info.Class = inst.class
	info.Scope = scope
	info.Aliases = inst.aliases
	info.factory = inst.simpleFactory

	// info.OnNew = inst.fnNew
	// info.OnInject = inst.fnInject
	// info.OnInit = inst.fnInit
	// info.OnDestroy = inst.fnDestroy

	inst.Reset()
	return info, nil
}

func (inst *comInfoBuilder) CreateTo(cb application.ConfigBuilder) error {
	info, err := inst.Create()
	if err != nil {
		return err
	}
	cb.AddComponent(info)
	return nil
}

func (inst *comInfoBuilder) parseScope(str string) (application.ComponentScope, error) {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	if str == "" {
		return application.ScopeSingleton, nil

	} else if str == "prototype" {
		return application.ScopePrototype, nil

	} else if str == "singleton" {
		return application.ScopeSingleton, nil
	}
	return 0, errors.New("bad component scope value:" + str)
}

func (inst *comInfoBuilder) parseAliases(str string) []string {

	const sp1 = " "
	const sp2 = ","
	str = strings.ReplaceAll(str, sp1, sp2)
	str = strings.ReplaceAll(str, "\t", sp2)
	list1 := strings.Split(str, sp2)
	list2 := []string{}

	for index := range list1 {
		item := list1[index]
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		list2 = append(list2, item)
	}

	return list2
}
