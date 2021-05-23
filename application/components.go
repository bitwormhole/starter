package application

import (
	"io"

	"github.com/bitwormhole/starter/lang"
)

// ComponentScope 枚举表示组件的作用域
type ComponentScope uint32

const (
	// ScopeMin 是作用域的最小值
	ScopeMin ComponentScope = 0 // 最小

	// ScopeSingleton 表示单例模式
	ScopeSingleton ComponentScope = 1

	// ScopeContext 表示上下文模式
	ScopeContext ComponentScope = 2

	// ScopePrototype 表示原型模式
	ScopePrototype ComponentScope = 3

	// ScopeMax 是作用域的最大值
	ScopeMax ComponentScope = 4 // 最大
)

// ComponentInstance  一个具体的组件的实例的引用
type ComponentInstance interface {
	Get() lang.Object
	IsLoaded() bool
	Inject(context RuntimeContext) error
	Init() error
	Destroy() error
}

// ComponentFactory 一个组件的工厂
type ComponentFactory interface {
	NewInstance() ComponentInstance
}

// ComponentInfo 一个组件的配置
type ComponentInfo interface {
	GetID() string
	GetClass() string
	GetAliases() []string
	GetClasses() []string
	GetScope() ComponentScope
	GetFactory() ComponentFactory
	GetPrototype() lang.Object

	IsTypeOf(typeName string) bool
	IsNameOf(alias string) bool
}

// ComponentHolder 一个具体的组件的代理
type ComponentHolder interface {
	GetInstance() ComponentInstance
	IsOriginalName(name string) bool
	GetInfo() ComponentInfo
	GetPrototype() lang.Object
	GetContext() RuntimeContext
	MakeChild(context RuntimeContext) ComponentHolder
}

type ComponentHolderFilter func(name string, holder ComponentHolder) bool

// ComponentLoader 用于加载组件的实例
type ComponentLoader interface {
	io.Closer
	Load(h ComponentHolder) (lang.Object, error)
	LoadAll(h []ComponentHolder) ([]lang.Object, error)
	GetReleasePool() lang.ReleasePool
	GetContext() RuntimeContext
}

// Components 接口表示一个组件的集合
type Components interface {
	// ids
	GetComponentNameList(includeAliases bool) []string

	// getters
	GetComponent(selector string) (ComponentHolder, error)
	GetComponents(selector string) []ComponentHolder
	GetComponentsByFilter(f ComponentHolderFilter) []ComponentHolder

	// export & import
	Export(map[string]ComponentHolder) map[string]ComponentHolder
	Import(map[string]ComponentHolder)
}
