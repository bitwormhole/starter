package application

import "github.com/bitwormhole/starter/collection"

// Configuration 表示应用程序配置
type Configuration interface {
	GetLoader() ContextLoader
	GetComponents() []ComponentInfo
	GetResources() collection.Resources
	GetEnvironment() collection.Environment
	IsEnableLoadPropertiesFromArguments() bool
}

//  ContextLoader 用于加载进程上下文
type ContextLoader interface {
	Load(config Configuration, args []string) (Context, error)
}

// ConfigBuilder 表示应用程序配置
type ConfigBuilder interface {
	AddComponent(info ComponentInfo)
	SetResources(res collection.Resources)

	SetEnableLoadPropertiesFromArguments(enable bool)
	IsEnableLoadPropertiesFromArguments() bool

	Create() Configuration
}
