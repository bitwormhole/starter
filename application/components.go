package application

import (
	"github.com/bitwormhole/starter/lang"
)

// ComponentScope 枚举表示组件的作用域
type ComponentScope uint32

// ComponentState  枚举表示组件的状态
type ComponentState uint32

const (
	ScopeMin       ComponentScope = 0 // ScopeMin 是作用域的最小值
	ScopeSingleton ComponentScope = 1 // ScopeSingleton 表示单例模式
	ScopeContext   ComponentScope = 2 // ScopeContext 表示上下文模式
	ScopePrototype ComponentScope = 3 // ScopePrototype 表示原型模式
	ScopeMax       ComponentScope = 4 // ScopeMax 是作用域的最大值
)

const (
	StateZero       ComponentState = 0 // 新建
	StateInjected   ComponentState = 1 // 已执行 injectMethod
	StateInitialled ComponentState = 2 // 已执行 initMethod
	StateReady      ComponentState = 3 // 正常可用
	StateDestroyed  ComponentState = 4 // 已执行 destroyMethod
)

// ComponentInstance  一个具体的组件的实例的引用
type ComponentInstance interface {
	Factory() ComponentFactory
	Get() lang.Object
	State() ComponentState
	Inject(context InstanceContext) error
	Init() error
	Destroy() error
}

// ComponentFactory 一个组件的工厂
type ComponentFactory interface {
	GetPrototype() lang.Object
	NewInstance() ComponentInstance
	AfterService() ComponentAfterService
}

// ComponentAfterService 是组件出厂后的服务
type ComponentAfterService interface {
	Inject(instance ComponentInstance, context InstanceContext) error
	Init(instance ComponentInstance) error
	Destroy(instance ComponentInstance) error
}

// InstanceContext  组件实例的上下文
type InstanceContext interface {
	Context() Context
	Pool() lang.ReleasePool

	// AddInstance(instance ComponentInstance)

	GetComponent(selector string) (lang.Object, error)
	GetComponents(selector string) ([]lang.Object, error)
	GetComponentsByFilter(selector string, f ComponentHolderFilter) ([]lang.Object, error)

	GetInt(selector string) (int, error)
	GetInt16(selector string) (int16, error)
	GetInt32(selector string) (int32, error)
	GetInt64(selector string) (int64, error)
	GetBool(selector string) (bool, error)
	GetString(selector string) (string, error)
}

// ComponentInfo 一个组件的配置
type ComponentInfo interface {
	GetID() string
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
	GetContext() Context
	MakeChild(context Context) ComponentHolder
}

// ComponentGroup 是holder的分组
type ComponentGroup interface {
	Size() int
	ListAll() []ComponentHolder
	ListWithFilter(f ComponentHolderFilter) []ComponentHolder
	Name() string
}

// ComponentGroupManager 用于管理holder的分组
type ComponentGroupManager interface {
	GetGroup(selector string) ComponentGroup
	Reload() error
}

// ComponentHolderFilter 是组件过滤器的函数签名
type ComponentHolderFilter func(name string, holder ComponentHolder) bool

// // ComponentLoader 用于加载组件的实例
// type ComponentLoader interface {
// 	OpenLoading(context Context) (ComponentLoading, error)
// }

// // ComponentLoading  表示加载组件的会话
// type ComponentLoading interface {
// 	io.Closer
// 	lang.ErrorHandler

// 	Pool() lang.ReleasePool
// 	Context() Context

// 	Load(h ComponentHolder) (lang.Object, error)
// 	LoadAll(h []ComponentHolder) ([]lang.Object, error)
// }

// Components 接口表示一个组件的集合
type Components interface {
	// ids
	GetComponentNameList(includeAliases bool) []string

	// finders
	FindComponent(selector string) (ComponentHolder, error)
	FindComponents(selector string) []ComponentHolder
	FindComponentsWithFilter(selector string, f ComponentHolderFilter) []ComponentHolder

	GroupManager() ComponentGroupManager

	// export & import
	Export(map[string]ComponentHolder) map[string]ComponentHolder
	Import(map[string]ComponentHolder)
}
