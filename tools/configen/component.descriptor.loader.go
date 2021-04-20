package configen

import (
	"strings"

	"github.com/bitwormhole/starter/collection"
)

type componentDescriptorLoader struct {
	context *configenContext
}

type componentDescriptorLoaderProps struct {
	prefix     string
	properties collection.Properties
}

////////////////////////////////////////////////////////////////////////////////
// impl componentDescriptorLoaderProps

func (inst *componentDescriptorLoaderProps) get(name string) string {
	key := inst.prefix + name
	return inst.properties.GetProperty(key, "")
}

////////////////////////////////////////////////////////////////////////////////
// impl componentDescriptorLoader

func (inst *componentDescriptorLoader) load() error {

	com_desc_list := []*componentDescriptor{}
	com_name_list := []string{}
	props := inst.context.properties
	src := props.Export(nil)

	for key := range src {
		com_name, ok := inst.parseComNameFromKey(key)
		if ok {
			com_name_list = append(com_name_list, com_name)
		}
	}

	for index := range com_name_list {
		name := com_name_list[index]
		com_desc, err := inst.loadComponentDescriptor(name, props)
		if err != nil {
			return err
		}
		com_desc_list = append(com_desc_list, com_desc)
	}

	inst.context.components = com_desc_list
	return nil
}

func (inst *componentDescriptorLoader) loadComponentDescriptor(name string, props collection.Properties) (*componentDescriptor, error) {

	cd := &componentDescriptor{}
	prefix := "component." + name + "."
	p := &componentDescriptorLoaderProps{prefix: prefix, properties: props}

	// required
	cd.descriptorName = name
	cd.typeName = p.get("type")

	// optional
	cd.id = p.get("id")
	cd.class = p.get("class")
	cd.aliases = p.get("aliases")
	cd.initMethod = p.get("initMethod")
	cd.destroyMethod = p.get("destroyMethod")
	cd.scope = p.get("scope")
	cd.inject = p.get("inject")

	return cd, nil
}

func (inst *componentDescriptorLoader) parseComNameFromKey(key string) (string, bool) {
	prefix := "component."
	suffix := ".type"
	if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
		name := key
		name = strings.TrimPrefix(name, prefix)
		name = strings.TrimSuffix(name, suffix)
		return name, true
	}
	return "", false
}
