package config

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

// InjectionSelector 代表一个注射选择器
type InjectionSelector interface {
	GetOne(context application.InstanceContext) lang.Object
	GetList(context application.InstanceContext) []lang.Object

	GetFloat64(context application.InstanceContext) float64
	GetString(context application.InstanceContext) string
	GetBool(context application.InstanceContext) bool
	GetInt(context application.InstanceContext) int
}

// NewInjectionSelector 新建一个注射选择器
func NewInjectionSelector(selector string, filter application.ComponentHolderFilter) InjectionSelector {
	inst := &innerInjectionSelector{}
	return inst.init(selector, filter)
}

////////////////////////////////////////////////////////////////////////////////

type innerInjectionSelector struct {
	selector string
	filter   application.ComponentHolderFilter
	ids      []string // cache for com.ids
}

func (inst *innerInjectionSelector) init(selector string, filter application.ComponentHolderFilter) InjectionSelector {
	inst.selector = selector
	inst.filter = filter
	return inst
}

func (inst *innerInjectionSelector) GetOne(context application.InstanceContext) lang.Object {
	o, err := context.GetComponent(inst.selector)
	if err != nil {
		context.HandleError(err)
		return nil
	}
	return o
}

func (inst *innerInjectionSelector) GetList(context application.InstanceContext) []lang.Object {
	list, err := context.GetComponentsByFilter(inst.selector, inst.filter)
	if err != nil {
		context.HandleError(err)
		return []lang.Object{}
	}
	return list
}

func (inst *innerInjectionSelector) GetString(context application.InstanceContext) string {
	value, err := context.GetString(inst.selector)
	if err != nil {
		context.HandleError(err)
		return ""
	}
	return value
}

func (inst *innerInjectionSelector) GetBool(context application.InstanceContext) bool {
	value, err := context.GetBool(inst.selector)
	if err != nil {
		context.HandleError(err)
		return false
	}
	return value
}

func (inst *innerInjectionSelector) GetInt(context application.InstanceContext) int {
	value, err := context.GetInt(inst.selector)
	if err != nil {
		context.HandleError(err)
		return 0
	}
	return value
}

func (inst *innerInjectionSelector) GetFloat64(context application.InstanceContext) float64 {
	value, err := context.GetFloat64(inst.selector)
	if err != nil {
		context.HandleError(err)
		return 0
	}
	return value
}
