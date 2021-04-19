package application

import "github.com/bitwormhole/starter/lang"

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

	IsTypeOf(typeName string) bool
	IsNameOf(alias string) bool
}

// ComponentHolder 一个具体的组件的代理
type ComponentHolder interface {
	GetInstance() ComponentInstance
	IsOriginalName(name string) bool
	GetInfo() ComponentInfo
	GetContext() RuntimeContext
	MakeChild(context RuntimeContext) ComponentHolder
}

// Components 接口表示一个组件的集合
type Components interface {
	GetComponent(name string) (lang.Object, error)
	GetComponentByClass(classSelector string) (lang.Object, error)
	GetComponentsByClass(classSelector string) []lang.Object
	GetComponentNameList(includeAliases bool) []string
	////
	Export(map[string]ComponentHolder) map[string]ComponentHolder
	Import(map[string]ComponentHolder)
}
