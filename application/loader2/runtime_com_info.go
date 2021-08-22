package loader2

import (
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

type runtimeComInfo struct {
	id        string
	aliases   []string
	classes   []string
	prototype lang.Object
	scope     application.ComponentScope
	// afterService application.ComponentAfterService
	innerFactory application.ComponentFactory
	outerfactory application.ComponentFactory
}

func (inst *runtimeComInfo) _Impl() (application.ComponentInfo, application.ComponentFactory) {
	return inst, inst
}

func (inst *runtimeComInfo) IsTypeOf(typeName string) bool {
	typeName = strings.TrimSpace(typeName)
	if typeName == "" {
		return false
	}
	list := inst.classes
	if list == nil {
		return false
	}
	for _, item := range list {
		if typeName == item {
			return true
		}
	}
	return false
}

func (inst *runtimeComInfo) IsNameOf(alias string) bool {

	alias = strings.TrimSpace(alias)
	if alias == "" {
		return false
	}

	if alias == inst.id {
		return true
	}

	list := inst.aliases
	if list == nil {
		return false
	}

	for _, item := range list {
		if alias == item {
			return true
		}
	}

	return false
}

func (inst *runtimeComInfo) GetAliases() []string {
	src := inst.aliases
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

func (inst *runtimeComInfo) GetClasses() []string {
	src := inst.classes
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

func (inst *runtimeComInfo) GetScope() application.ComponentScope {
	return inst.scope
}

func (inst *runtimeComInfo) GetID() string {
	return inst.id
}

func (inst *runtimeComInfo) GetFactory() application.ComponentFactory {
	return inst
}

func (inst *runtimeComInfo) GetPrototype() lang.Object {
	return inst.prototype
}

func (inst *runtimeComInfo) AfterService() application.ComponentAfterService {
	return inst.innerFactory.AfterService()
}

func (inst *runtimeComInfo) NewInstance() application.ComponentInstance {
	i := inst.innerFactory.NewInstance()
	wrapper := &comInstanceWrapper{}
	return wrapper.init(i, inst)
}

////////////////////////////////////////////////////////////////////////////////

type runtimeComInfoBuilder struct {
	defaultID string
	info      application.ComponentInfo
}

func (inst *runtimeComInfoBuilder) init(info application.ComponentInfo) {
	inst.info = info
}

func (inst *runtimeComInfoBuilder) setDefaultID(id string) {
	inst.defaultID = id
}

func (inst *runtimeComInfoBuilder) Create() (application.ComponentInfo, error) {

	src := inst.info
	factoryIn := src.GetFactory()
	classes := src.GetClasses()
	aliases := src.GetAliases()
	id := strings.TrimSpace(src.GetID())
	scope := src.GetScope()

	// for default values

	if id == "" {
		id = strings.TrimSpace(inst.defaultID)
	}

	if classes == nil {
		classes = []string{}
	}

	if aliases == nil {
		aliases = []string{}
	}

	if (scope <= application.ScopeMin) || (scope >= application.ScopeMax) {
		scope = application.ScopeSingleton
	}

	// make
	info := &runtimeComInfo{}
	//	info.afterService = after
	info.aliases = aliases
	info.classes = classes
	info.id = id
	info.prototype = factoryIn.GetPrototype()
	info.scope = scope
	info.innerFactory = factoryIn
	info.outerfactory = info
	return info, nil
}
