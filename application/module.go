package application

import (
	"strconv"

	"github.com/bitwormhole/starter/collection"
)

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

////////////////////////////////////////////////////////////////////////////////

// DefineModule 定义一个模块
type DefineModule struct {
	Name         string
	Version      string
	Revision     int
	Resources    collection.Resources
	Dependencies []Module
	OnMount      OnMountFunc
}

func (inst *DefineModule) _Impl() Module {
	return inst
}

// AddDependency 添加一个依赖
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

// GetDependencies 返回依赖的其它模块
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

// GetResources 取模块的资源
func (inst *DefineModule) GetResources() collection.Resources {
	return inst.Resources
}

// GetName 取模块名称
func (inst *DefineModule) GetName() string {
	return inst.Name
}

// GetRevision 取模块修订编号
func (inst *DefineModule) GetRevision() int {
	return inst.Revision
}

// GetVersion 取模块版本
func (inst *DefineModule) GetVersion() string {
	return inst.Version
}

// Apply 向 ConfigBuilder 注册本模块中包含的组件，并注入默认配置
func (inst *DefineModule) Apply(cb ConfigBuilder) error {
	return inst.OnMount(cb)
}

////////////////////////////////////////////////////////////////////////////////

// ModuleBuilder 用于创建Module对象
type ModuleBuilder struct {
	name     string
	version  string
	revision int
	deps     []Module
	res      collection.Resources
	onMount  OnMountFunc
}

// Name 设置模块的名称
func (inst *ModuleBuilder) Name(name string) *ModuleBuilder {
	inst.name = name
	return inst
}

// Version 设置模块的版本
func (inst *ModuleBuilder) Version(version string) *ModuleBuilder {
	inst.version = version
	return inst
}

// Revision 设置模块的修订编号
func (inst *ModuleBuilder) Revision(revision int) *ModuleBuilder {
	inst.revision = revision
	return inst
}

// Resources 设置跟模块绑定的资源
func (inst *ModuleBuilder) Resources(resources collection.Resources) *ModuleBuilder {
	inst.res = resources
	return inst
}

// Dependencies 添加一组依赖
func (inst *ModuleBuilder) Dependencies(mods []Module) *ModuleBuilder {
	if mods == nil {
		return inst
	}
	for _, mod := range mods {
		inst.Dependency(mod)
	}
	return inst
}

// Dependency 添加一个依赖
func (inst *ModuleBuilder) Dependency(mod Module) *ModuleBuilder {
	if mod == nil {
		return inst
	}
	list := inst.getDependencyList(true)
	list = append(list, mod)
	inst.deps = list
	return inst
}

// OnMount 设置配置模块的入口函数
func (inst *ModuleBuilder) OnMount(fn OnMountFunc) *ModuleBuilder {
	inst.onMount = fn
	return inst
}

func (inst *ModuleBuilder) getDependencyList(enableInit bool) []Module {
	list := inst.deps
	if list == nil && enableInit {
		list = make([]Module, 0)
		inst.deps = list
	}
	return list
}

var theModuleBuilderIndexCount = 1

// Create 创建模块
func (inst *ModuleBuilder) Create() Module {

	name := inst.name
	ver := inst.version
	rev := inst.revision
	deps := inst.getDependencyList(true)
	res := inst.res
	onMount := inst.onMount
	index := theModuleBuilderIndexCount

	theModuleBuilderIndexCount++

	if name == "" {
		name = "unnamed-module-" + strconv.Itoa(index)
	}

	if ver == "" {
		ver = "v1"
	}

	if res == nil {
		res = collection.CreateResources()
	}

	if onMount == nil {
		onMount = func(cb ConfigBuilder) error { return nil }
	}

	return &DefineModule{
		Name:         name,
		Version:      ver,
		Revision:     rev,
		Dependencies: deps,
		Resources:    res,
		OnMount:      onMount,
	}
}
