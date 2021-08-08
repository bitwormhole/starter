package application

import "github.com/bitwormhole/starter/collection"

// Module 表示一个可导入的模块
type Module interface {
	GetName() string
	GetVersion() string
	GetRevision() int
	GetResources() collection.Resources
	GetDependencies() []Module
	Apply(cb ConfigBuilder) error
}

// OnMountFunc 是模块挂载函数的签名
type OnMountFunc func(cb ConfigBuilder) error

// DefineModule 定义一个模块
type DefineModule struct {
	Name         string
	Version      string
	Revision     int
	Resources    collection.Resources
	Dependencies []Module
	OnMount      OnMountFunc
}

func (inst *DefineModule) __impl__() Module {
	return inst
}

func (inst *DefineModule) AddDependency(mod Module) {
	if mod == nil {
		return
	}
	list := inst.Dependencies
	if list == nil {
		list = make([]Module, 1)
		list[0] = mod
	} else {
		list = append(list, mod)
	}
	inst.Dependencies = list
}

func (inst *DefineModule) GetDependencies() []Module {
	src := inst.Dependencies
	dst := make([]Module, 0)
	if src == nil {
		return dst
	}
	for index := range src {
		mod := src[index]
		if mod != nil {
			dst = append(dst, mod)
		}
	}
	return dst
}

func (inst *DefineModule) GetResources() collection.Resources {
	return inst.Resources
}

func (inst *DefineModule) GetName() string {
	return inst.Name
}

func (inst *DefineModule) GetRevision() int {
	return inst.Revision
}

func (inst *DefineModule) GetVersion() string {
	return inst.Version
}

func (inst *DefineModule) Apply(cb ConfigBuilder) error {
	return inst.OnMount(cb)
}
