package config

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/loader"
	"github.com/bitwormhole/starter/collection"
)

////////////////////////////////////////////////////////////////////////////////

// AppConfig 提供一个简易的 Configuration 实现
type appConfig struct {
	// implements Configuration
	components                   []application.ComponentInfo
	resources                    collection.Resources
	enableLoadPropertiesFromArgs bool
}

func (inst *appConfig) init() application.ConfigBuilder {
	inst.components = []application.ComponentInfo{}
	inst.enableLoadPropertiesFromArgs = true
	return inst
}

func (inst *appConfig) getComList(create bool) []application.ComponentInfo {
	list := inst.components
	if (create) && (list == nil) {
		list = make([]application.ComponentInfo, 0)
		inst.components = list
	}
	return list
}

// GetComponents 返回组件的注册信息
func (inst *appConfig) GetComponents() []application.ComponentInfo {
	return inst.getComList(true)
}

// Create 用于创建配置
func (inst *appConfig) Create() application.Configuration {
	return inst
}

// GetEnvironment 用于env
func (inst *appConfig) GetEnvironment() collection.Environment {
	return nil
}

func (inst *appConfig) IsEnableLoadPropertiesFromArguments() bool {
	return inst.enableLoadPropertiesFromArgs
}

func (inst *appConfig) SetEnableLoadPropertiesFromArguments(enable bool) {
	inst.enableLoadPropertiesFromArgs = enable
}

// AddComponent 注册一个组件
func (inst *appConfig) AddComponent(info application.ComponentInfo) {
	list := inst.getComList(true)
	inst.components = append(list, info)
}

// GetLoader 返回加载器
func (inst *appConfig) GetLoader() application.ContextLoader {
	return &loader.RuntimeContextLoader{}
}

// GetBuilder 返回构建器
func (inst *appConfig) GetBuilder() application.ConfigBuilder {
	return inst
}

// SetResources 用于配置上下文的资源文件夹
func (inst *appConfig) SetResFS(fs *embed.FS, prefix string) {
	inst.resources = &simpleEmbedResFS{
		fs:     fs,
		prefix: prefix,
	}
}

func (inst *appConfig) SetResources(res collection.Resources) {
	inst.resources = res
}

// GetResources 用于获取上下文的资源文件夹
func (inst *appConfig) GetResources() collection.Resources {
	return inst.resources
}

////////////////////////////////////////////////////////////////////////////////
// Builder

func NewBuilder() application.ConfigBuilder {
	cfg := &appConfig{}
	cb := cfg.init()
	return cb
}

func NewBuilderRes(res collection.Resources) application.ConfigBuilder {
	cfg := &appConfig{}
	cb := cfg.init()
	cb.SetResources(res)
	return cb
}

func NewBuilderFS(fs *embed.FS, prefix string) application.ConfigBuilder {
	cfg := &appConfig{}
	cb := cfg.init()
	cfg.SetResFS(fs, prefix)
	return cb
}

////////////////////////////////////////////////////////////////////////////////
// EOF
