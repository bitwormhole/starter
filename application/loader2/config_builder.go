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
	hError     lang.ErrorHandler
	components []application.ComponentInfo

	resources         collection.Resources
	propertiesDefault collection.Properties
	propertiesFinal   collection.Properties
	attributes        collection.Attributes

	enableLoadPropertiesFromArgs bool
}

func (inst *configBuilderV2) init() application.ConfigBuilder {
	inst.enableLoadPropertiesFromArgs = true
	inst.propertiesDefault = collection.CreateProperties()
	inst.propertiesFinal = collection.CreateProperties()
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
	inst.AddProperties2(src, nil)
}

func (inst *configBuilderV2) AddProperties2(def, final collection.Properties) {

	if def != nil {
		dst := inst.propertiesDefault
		if dst == nil {
			dst = collection.CreateProperties()
			inst.propertiesDefault = dst
		}
		list := def.Export(nil)
		dst.Import(list)
	}

	if final != nil {
		dst := inst.propertiesFinal
		if dst == nil {
			dst = collection.CreateProperties()
			inst.propertiesFinal = dst
		}
		list := final.Export(nil)
		dst.Import(list)
	}
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
	inst.AddResources(res)
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
	return inst.propertiesDefault
}

func (inst *configBuilderV2) FinalProperties() collection.Properties {
	return inst.propertiesFinal
}

func (inst *configBuilderV2) Create() application.Configuration {
	cfg := &configuration{}
	return cfg.init(inst)
}
