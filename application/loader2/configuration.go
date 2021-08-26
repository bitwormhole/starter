package loader2

import (
	"os"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

type configuration struct {
	cb     *configBuilderV2
	loader application.ContextLoader
	env    collection.Environment
}

func (inst *configuration) init(cb *configBuilderV2) application.Configuration {
	inst.cb = cb
	inst.loader = &loader{}
	inst.env = inst.loadEnv()
	return inst
}

func (inst *configuration) loadEnv() collection.Environment {
	env := collection.CreateEnvironment()
	list := os.Environ()
	for _, str := range list {
		kv := strings.SplitN(str, "=", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])
			env.SetEnv(key, val)
		}
	}
	return env
}

func (inst *configuration) GetLoader() application.ContextLoader {
	return inst.loader
}

func (inst *configuration) GetComponents() []application.ComponentInfo {
	return inst.cb.components

}

func (inst *configuration) GetResources() collection.Resources {
	return inst.cb.resources
}

func (inst *configuration) GetAttributes() collection.Attributes {
	return inst.cb.attributes
}

func (inst *configuration) GetEnvironment() collection.Environment {
	return inst.env
}

func (inst *configuration) GetDefaultProperties() collection.Properties {
	return inst.cb.properties
}

func (inst *configuration) IsEnableLoadPropertiesFromArguments() bool {
	return inst.cb.enableLoadPropertiesFromArgs
}