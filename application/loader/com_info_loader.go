package loader

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// componentInfoPreprocessor class
type componentInfoPreprocessor struct {
}

func (inst *componentInfoPreprocessor) prepare(info1 application.ComponentInfo, index int) (application.ComponentInfo, error) {

	id := info1.GetID()
	aliases := info1.GetAliases()
	classes := info1.GetClasses()
	classDef := info1.GetClass()
	scope := info1.GetScope()
	factory := info1.GetFactory()

	id = strings.TrimSpace(id)

	if id == "" {
		id = fmt.Sprint("com", index, "_", strings.TrimSpace(classDef))
	}

	if factory == nil {
		return nil, errors.New("no factory for component named:" + id)
	}

	if scope <= application.ScopeMin || application.ScopeMax <= scope {
		scope = application.ScopeSingleton
	}

	classes = inst.normalizeClasses(classes, classDef)
	aliases = inst.normalizeAliases(aliases, id)
	factory = createComInstanceWithStateFactory(factory)

	info2 := &preparedComponentInfo{}
	info2.id = id
	info2.classDefine = classDef
	info2.aliases = aliases
	info2.classes = classes
	info2.factory = factory
	info2.scope = scope
	info2.prototype = factory.NewInstance().Get()
	return info2, nil
}

func (inst *componentInfoPreprocessor) map2array(table map[string]bool) []string {
	array := make([]string, 0)
	for name := range table {
		if name == "" {
			continue
		}
		if table[name] {
			array = append(array, name)
		}
	}
	return array
}

func (inst *componentInfoPreprocessor) normalizeClasses(classes []string, classDef string) []string {
	set := make(map[string]bool)
	array := strings.Split(classDef, " ")
	for index := range array {
		item := strings.TrimSpace(array[index])
		set[item] = true
	}
	if classes != nil {
		for index := range classes {
			item := strings.TrimSpace(classes[index])
			set[item] = true
		}
	}
	return inst.map2array(set)
}

func (inst *componentInfoPreprocessor) normalizeAliases(aliases []string, id string) []string {
	set := make(map[string]bool)
	if aliases != nil {
		for index := range aliases {
			item := strings.TrimSpace(aliases[index])
			set[item] = true
		}
	}
	set[id] = true
	return inst.map2array(set)
}

////////////////////////////////////////////////////////////////////////////////
// preparedComponentInfo class
type preparedComponentInfo struct {
	id          string
	classDefine string
	aliases     []string
	classes     []string

	scope     application.ComponentScope
	factory   application.ComponentFactory
	prototype lang.Object
}

func (inst *preparedComponentInfo) GetID() string {
	return inst.id
}

func (inst *preparedComponentInfo) GetClass() string {
	return inst.classDefine
}

func (inst *preparedComponentInfo) GetAliases() []string {
	return inst.aliases
}

func (inst *preparedComponentInfo) GetClasses() []string {
	return inst.classes
}

func (inst *preparedComponentInfo) GetScope() application.ComponentScope {
	return inst.scope
}

func (inst *preparedComponentInfo) GetFactory() application.ComponentFactory {
	return inst.factory
}

func (inst *preparedComponentInfo) GetPrototype() lang.Object {
	return inst.prototype
}

func (inst *preparedComponentInfo) IsTypeOf(typeName string) bool {
	classes := inst.classes
	if classes == nil {
		return false
	}
	for index := range classes {
		item := classes[index]
		if typeName == item {
			return true
		}
	}
	return false
}

func (inst *preparedComponentInfo) IsNameOf(name string) bool {
	if name == "" {
		return false
	}
	if inst.id == name {
		return true
	}
	aliases := inst.aliases
	if aliases == nil {
		return false
	}
	for index := range aliases {
		alias := aliases[index]
		if alias == name {
			return true
		}
	}
	return false
}
