package config

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/runtime"
	"github.com/bitwormhole/starter/collection"
)

////////////////////////////////////////////////////////////////////////////////

// AppConfig 提供一个简易的 Configuration 实现
type AppConfig struct {
	// implements Configuration
	components []application.ComponentInfo
	resources  collection.Resources
}

func (inst *AppConfig) getComList(create bool) []application.ComponentInfo {
	list := inst.components
	if (create) && (list == nil) {
		list = make([]application.ComponentInfo, 0)
		inst.components = list
	}
	return list
}

// GetComponents 返回组件的注册信息
func (inst *AppConfig) GetComponents() []application.ComponentInfo {
	return inst.getComList(true)
}

// Create 用于创建配置
func (inst *AppConfig) Create() application.Configuration {
	return inst
}

// GetEnvironment 用于env
func (inst *AppConfig) GetEnvironment() collection.Environment {
	return nil
}

// AddComponent 注册一个组件
func (inst *AppConfig) AddComponent(info application.ComponentInfo) {
	list := inst.getComList(true)
	inst.components = append(list, info)
}

// GetLoader 返回加载器
func (inst *AppConfig) GetLoader() application.ContextLoader {
	return &runtime.RuntimeContextLoader{}
}

// GetBuilder 返回构建器
func (inst *AppConfig) GetBuilder() application.ConfigBuilder {
	return inst
}

// SetResources 用于配置上下文的资源文件夹
func (inst *AppConfig) SetResFS(fs *embed.FS, prefix string) {
	inst.resources = &simpleEmbedResFS{
		fs:     fs,
		prefix: prefix,
	}
}

func (inst *AppConfig) SetResources(res collection.Resources) {
	inst.resources = res
}

// GetResources 用于获取上下文的资源文件夹
func (inst *AppConfig) GetResources() collection.Resources {
	return inst.resources
}

////////////////////////////////////////////////////////////////////////////////
// Builder

func NewBuilderFS(fs *embed.FS, prefix string) application.ConfigBuilder {
	cfg := &AppConfig{}
	cfg.SetResFS(fs, prefix)
	return cfg
}

func NewBuilder() application.ConfigBuilder {
	return &AppConfig{}
}

////////////////////////////////////////////////////////////////////////////////
// EOF
