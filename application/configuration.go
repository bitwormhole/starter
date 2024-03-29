package application

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// Configuration 表示应用程序配置
type Configuration interface {
	GetLoader() ContextLoader
	GetModules() []Module
	GetComponents() []ComponentInfo
	GetResources() collection.Resources
	GetAttributes() collection.Attributes
	GetEnvironment() collection.Environment
	GetDefaultProperties() collection.Properties
	GetFinalProperties() collection.Properties
	GetErrorHandler() lang.ErrorHandler
	IsEnableLoadPropertiesFromArguments() bool
}

// ContextLoader 用于加载进程上下文
type ContextLoader interface {
	Load(config Configuration, args []string) (Context, error)
}

// ConfigBuilder 表示应用程序配置
type ConfigBuilder interface {
	AddComponent(info ComponentInfo)
	AddResources(res collection.Resources)
	AddProperties(p collection.Properties)
	AddProperties2(theDefault, theFinal collection.Properties)

	SetModules(mods []Module)

	SetResources(res collection.Resources)
	SetAttribute(name string, value interface{})

	SetErrorHandler(h lang.ErrorHandler)

	SetEnableLoadPropertiesFromArguments(enable bool)
	IsEnableLoadPropertiesFromArguments() bool

	DefaultProperties() collection.Properties
	FinalProperties() collection.Properties

	Create() Configuration
}
