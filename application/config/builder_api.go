package config

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// 定义 comInfo 的函数签名

// OnNew 新建组件对象
type OnNew func() lang.Object

// OnInject 向组件对象注入依赖
type OnInject func(obj lang.Object, context application.InstanceContext) error

// OnInit 初始化组件对象
type OnInit func(obj lang.Object) error

// OnDestroy 销毁组件对象
type OnDestroy func(obj lang.Object) error

////////////////////////////////////////////////////////////////////////////////

// ComponentInfoBuilder 接口用于构建组件信息
type ComponentInfoBuilder interface {
	ID(id string) ComponentInfoBuilder
	Class(class string) ComponentInfoBuilder
	Aliases(aliases string) ComponentInfoBuilder
	Scope(scope string) ComponentInfoBuilder

	Factory(f application.ComponentFactory) ComponentInfoBuilder

	// // OnNew 方法【已废弃】
	// OnNew(fn OnNew) ComponentInfoBuilder
	// // OnInject 方法【已废弃】
	// OnInject(fn OnInject) ComponentInfoBuilder
	// // OnInit 方法【已废弃】
	// OnInit(fn OnInit) ComponentInfoBuilder
	// // OnDestroy 方法【已废弃】
	// OnDestroy(fn OnDestroy) ComponentInfoBuilder

	Next() ComponentInfoBuilder
	Create() (application.ComponentInfo, error)
	CreateTo(cb application.ConfigBuilder) error
}

////////////////////////////////////////////////////////////////////////////////

// ComInfo  开始登记组件信息
func ComInfo() ComponentInfoBuilder {
	return &comInfoBuilder{}
}
