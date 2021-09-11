package application

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// Initializer 是应用程序的启动器
type Initializer interface {
	Run()
	RunEx() (Runtime, error)

	SetErrorHandler(h lang.ErrorHandler) Initializer
	SetAttribute(name string, value interface{}) Initializer
	SetExitEnabled(enabled bool) Initializer
	SetPanicEnabled(enabled bool) Initializer

	UseResources(res collection.Resources) Initializer
	UseProperties(res collection.Properties) Initializer
	Use(module Module) Initializer
	UsePanic() Initializer
}

// Runtime 提供运行应用程序的更多选项
type Runtime interface {
	Context() Context
	Loop() error
	Exit() error
}
