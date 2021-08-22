package bootstrap

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/loader2"
	"github.com/bitwormhole/starter/collection"
)

// ConfigBuilder 创建默认的 application.ConfigBuilder
func ConfigBuilder() application.ConfigBuilder {
	return loader2.ConfigBuilder()
}

// NewBuilder 新建配置建造器
func NewBuilder() application.ConfigBuilder {
	return loader2.ConfigBuilder()
}

// NewBuilderRes 新建配置建造器，并附带指定的资源
func NewBuilderRes(res collection.Resources) application.ConfigBuilder {
	builder := loader2.ConfigBuilder()
	builder.SetResources(res)
	return builder
}
