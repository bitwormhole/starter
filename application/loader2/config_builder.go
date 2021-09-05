package loader2

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// ConfigBuilder 函数新建一个v2的application.ConfigBuilder对象
func ConfigBuilder() application.ConfigBuilder {
	cb := &configBuilderV2{}
	return cb.init()
}

////////////////////////////////////////////////////////////////////////////////

type configBuilderV2 struct {
	hError                       lang.ErrorHandler
	components                   []application.ComponentInfo
	resources                    collection.Resources
	properties                   collection.Properties
	attributes                   collection.Attributes
	enableLoadPropertiesFromArgs bool
}

func (inst *configBuilderV2) init() application.ConfigBuilder {
	inst.enableLoadPropertiesFromArgs = true
	inst.properties = collection.CreateProperties()
	inst.components = make([]application.ComponentInfo, 0)
	inst.attributes = collection.CreateAttributes()
	inst.resources = collection.CreateResources()
	return inst
}

func (inst *configBuilderV2) AddComponent(info application.ComponentInfo) {
	if info == nil {
		return
	}
	inst.components = append(inst.components, info)
}

func (inst *configBuilderV2) AddProperties(src collection.Properties) {
	if src == nil {
		return
	}
	dst := inst.properties
	if dst == nil {
		dst = collection.CreateProperties()
		inst.properties = dst
	}
	list := src.Export(nil)
	dst.Import(list)
}

func (inst *configBuilderV2) AddResources(src collection.Resources) {
	if src == nil {
		return
	}
	dst := inst.resources
	if dst == nil {
		dst = collection.CreateResources()
		inst.resources = dst
	}
	list := src.Export(nil)
	dst.Import(list, true)
}

func (inst *configBuilderV2) SetResources(res collection.Resources) {
	if res == nil {
		return
	}
	inst.resources = res
}

func (inst *configBuilderV2) SetAttribute(name string, value interface{}) {
	inst.attributes.SetAttribute(name, value)
}

func (inst *configBuilderV2) SetErrorHandler(h lang.ErrorHandler) {
	inst.hError = h
}

func (inst *configBuilderV2) SetEnableLoadPropertiesFromArguments(enable bool) {
	inst.enableLoadPropertiesFromArgs = enable
}

func (inst *configBuilderV2) IsEnableLoadPropertiesFromArguments() bool {
	return inst.enableLoadPropertiesFromArgs
}

func (inst *configBuilderV2) DefaultProperties() collection.Properties {
	return inst.properties
}

func (inst *configBuilderV2) Create() application.Configuration {
	cfg := &configuration{}
	return cfg.init(inst)
}
