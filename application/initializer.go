package application

import "github.com/bitwormhole/starter/lang"

// Initializer 是应用程序的启动器
type Initializer interface {

	// EmbedResources(fs *embed.FS, path string) Initializer
	// MountResources(res collection.Resources) Initializer

	SetErrorHandler(h lang.ErrorHandler) Initializer
	SetAttribute(name string, value interface{}) Initializer
	Use(module Module) Initializer
	UsePanic() Initializer
	Run()
	RunEx() (Runtime, error)
}

// Runtime 提供运行应用程序的更多选项
type Runtime interface {
	Context() Context
	Loop() error
	Exit() error
}
