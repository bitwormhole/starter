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

	GetComponent(selector string) (lang.Object, error)
	GetComponentList(selector string) ([]lang.Object, error)

	Injector() Injector
	ComponentLoader() ComponentLoader
}

type SimpleContext interface {
	collection.Atts
}
