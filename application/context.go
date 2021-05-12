package application

import (
	"io"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// Context 表示一个通用的上下文对象
type Context interface {
	GetComponents() Components
	GetReleasePool() collection.ReleasePool
	GetArguments() collection.Arguments
	GetAttributes() collection.Attributes
	GetEnvironment() collection.Environment
	GetProperties() collection.Properties
	GetParameters() collection.Parameters
	GetResources() collection.Resources
}

// RuntimeContext 是app的全局上下文
type RuntimeContext interface {
	Context

	// info
	GetURI() string
	GetApplicationName() string
	GetApplicationVersion() string
	GetStartupTimestamp() int64
	GetShutdownTimestamp() int64

	// helper
	SetErrorHandler(h lang.ErrorHandler)
	GetErrorHandler() lang.ErrorHandler
	NewChild() RuntimeContext
	OpenCreationContext(scope ComponentScope) CreationContext
}

// CreationContext 是构建时的上下文
type CreationContext interface {
	io.Closer
	GetScope() ComponentScope
	GetContext() RuntimeContext
}
