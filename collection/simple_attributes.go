package collection

import (
	"errors"

	"github.com/bitwormhole/starter/lang"
)

type SimpleAttributes struct {
	table map[string]lang.Object
}

func (inst *SimpleAttributes) _impl_() Attributes {
	return inst
}

func (inst *SimpleAttributes) GetAttribute(name string) lang.Object {
	table := inst.table
	if table == nil {
		return nil
	}
	return table[name]
}

func (inst *SimpleAttributes) GetAttributeRequired(name string) (lang.Object, error) {
	value := inst.GetAttribute(name)
	if value == nil {
		return nil, errors.New("no attr named:" + name)
	}
	return value, nil
}

func (inst *SimpleAttributes) SetAttribute(name string, value lang.Object) {
	table := inst.table
	if table == nil {
		table = make(map[string]lang.Object)
	}
	table[name] = value
	inst.table = table
}

func (inst *SimpleAttributes) Export(dst map[string]lang.Object) map[string]lang.Object {
	if dst == nil {
		dst = make(map[string]lang.Object)
	}
	src := inst.table
	if src == nil {
		return dst
	}
	for key := range src {
		val := src[key]
		dst[key] = val
	}
	return dst
}

func (inst *SimpleAttributes) Import(src map[string]lang.Object) {
	if src == nil {
		return
	}
	dst := inst.table
	if dst == nil {
		dst = make(map[string]lang.Object)
	}
	for key := range src {
		val := src[key]
		dst[key] = val
	}
	inst.table = dst
}

func CreateAttributes() Attributes {
	table := make(map[string]lang.Object)
	return &SimpleAttributes{table: table}
}
