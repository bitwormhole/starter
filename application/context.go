package application

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// ContextInfo 包含上下文的基本信息
type ContextInfo interface {

	// info
	GetURI() string
	GetApplicationName() string
	GetApplicationVersion() string
	GetStartupTimestamp() int64
	GetShutdownTimestamp() int64
}

// ContextCollections 提供一组getter，来获取context的各种集合
type ContextCollections interface {

	// GetReleasePool 取context的生命周期管理池
	GetReleasePool() lang.ReleasePool

	// GetComponents 取context组件管理器
	GetComponents() Components

	// GetResources 取context的资源管理器
	GetResources() collection.Resources

	GetArguments() collection.Arguments
	GetAttributes() collection.Attributes
	GetEnvironment() collection.Environment
	GetProperties() collection.Properties
	GetParameters() collection.Parameters
}

// Context 表示一个通用的上下文对象
type Context interface {
	ContextCollections
	ContextInfo

	// helper
	SetErrorHandler(h lang.ErrorHandler)
	GetErrorHandler() lang.ErrorHandler
	NewChild() Context

	GetComponent(selector string) (lang.Object, error)
	GetComponentList(selector string) ([]lang.Object, error)

	// 【准备添加】
	// GetComponentsByFilter(selector string, f ComponentHolderFilter) ([]lang.Object, error)

	// 【准备废弃】
	//Injector() Injector

	// 【准备废弃】
	// ComponentLoader() ComponentLoader
}

// SimpleContext 【已废弃】用“context.Context” & “lang.Context” 代替
type SimpleContext interface {
	collection.Atts
}
