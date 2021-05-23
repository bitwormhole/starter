package application

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

type ContextInfo interface {

	// info
	GetURI() string
	GetApplicationName() string
	GetApplicationVersion() string
	GetStartupTimestamp() int64
	GetShutdownTimestamp() int64
}

type ContextCollections interface {
	GetReleasePool() lang.ReleasePool
	GetComponents() Components

	GetArguments() collection.Arguments
	GetAttributes() collection.Attributes
	GetEnvironment() collection.Environment
	GetProperties() collection.Properties
	GetParameters() collection.Parameters
	GetResources() collection.Resources
}

// Context 表示一个通用的上下文对象
type Context interface {
	ContextCollections
	ContextInfo

	// helper
	SetErrorHandler(h lang.ErrorHandler)
	GetErrorHandler() lang.ErrorHandler
	NewChild() Context
	//	OpenCreationContext(scope ComponentScope) CreationContext

	FindComponent(selector string) (lang.Object, error)
	FindComponents(selector string) []lang.Object
	Injector() Injector
	InjectorScope(scope ComponentScope) Injector
}

// RuntimeContext 是app的全局上下文(alias for Context)
type RuntimeContext interface {
	Context
}
