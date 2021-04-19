package collection

import (
	"errors"

	"github.com/bitwormhole/starter/lang"
)

type SimpleAttributes struct {
	table map[string]lang.Object
}

func (inst *SimpleAttributes) GetAttribute(name string) (lang.Object, error) {
	table := inst.table
	if table == nil {
		return nil, errors.New("no attr named:" + name)
	}
	val := table[name]
	if val == nil {
		return nil, errors.New("no attr named:" + name)
	}
	return val, nil
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
