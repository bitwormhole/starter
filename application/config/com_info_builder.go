package config

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter/application"
)

type ComInfoBuilder struct {
	id      string
	class   string
	aliases string
	scope   string

	fnNew     OnNew
	fnInject  OnInject
	fnInit    OnInit
	fnDestroy OnDestroy
}

func (inst *ComInfoBuilder) ID(id string) *ComInfoBuilder {
	inst.id = id
	return inst
}

func (inst *ComInfoBuilder) Class(cls string) *ComInfoBuilder {
	inst.class = cls
	return inst
}

func (inst *ComInfoBuilder) Aliases(aliases string) *ComInfoBuilder {
	inst.aliases = aliases
	return inst
}

func (inst *ComInfoBuilder) Scope(scope string) *ComInfoBuilder {
	inst.scope = scope
	return inst
}

func (inst *ComInfoBuilder) OnNew(fn OnNew) *ComInfoBuilder {
	inst.fnNew = fn
	return inst
}

func (inst *ComInfoBuilder) OnInject(fn OnInject) *ComInfoBuilder {
	inst.fnInject = fn
	return inst
}

func (inst *ComInfoBuilder) OnInit(fn OnInit) *ComInfoBuilder {
	inst.fnInit = fn
	return inst
}

func (inst *ComInfoBuilder) OnDestroy(fn OnDestroy) *ComInfoBuilder {
	inst.fnDestroy = fn
	return inst
}

func (inst *ComInfoBuilder) Reset() {

	inst.id = ""
	inst.class = ""
	inst.scope = ""
	inst.aliases = ""

	inst.fnNew = nil
	inst.fnInject = nil
	inst.fnInit = nil
	inst.fnDestroy = nil
}

func (inst *ComInfoBuilder) Create() (*ComInfo, error) {

	scope, err := inst.parseScope(inst.scope)
	if err != nil {
		return nil, err
	}

	info := &ComInfo{}

	info.ID = inst.id
	info.Class = inst.class
	info.Scope = scope
	info.Aliases = inst.parseAliases(inst.aliases)

	info.OnNew = inst.fnNew
	info.OnInject = inst.fnInject
	info.OnInit = inst.fnInit
	info.OnDestroy = inst.fnDestroy

	inst.Reset()
	if info.OnNew == nil {
		return nil, errors.New("no func OnNew() for ComInfo")
	}
	return info, nil
}

func (inst *ComInfoBuilder) CreateTo(cb application.ConfigBuilder) error {
	info, err := inst.Create()
	if err != nil {
		return err
	}
	cb.AddComponent(info)
	return nil
}

func (inst *ComInfoBuilder) parseScope(str string) (application.ComponentScope, error) {
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

func (inst *ComInfoBuilder) parseAliases(str string) []string {

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
