package configen

import (
	"errors"
	"strings"
)

type componentBuildingInfoMaker struct {
	context *configenContext
}

func (inst *componentBuildingInfoMaker) make() error {
	src := inst.context.components
	dst := []*componentBuildingInfo{}
	for index := range src {
		item := src[index]
		info, err := inst.makeComInfo(item)
		if err != nil {
			return err
		}
		dst = append(dst, info)
	}
	inst.context.comBuildingInfoList = dst
	return nil
}

func (inst *componentBuildingInfoMaker) makeComInfo(desc *componentDescriptor) (*componentBuildingInfo, error) {

	info := &componentBuildingInfo{}

	info.id = desc.id
	info.aliases = inst.parseAliases(desc.aliases)
	info.classes = inst.parseClasses(desc.class)
	info.initMethod = desc.initMethod
	info.destroyMethod = desc.destroyMethod
	info.scope = desc.scope
	info.inject = desc.inject

	err := inst.parseTypeName(desc.typeName, info)
	if err != nil {
		return nil, err
	}

	importer, err := inst.context.importerManager.loadImporter(info.typePackageName)
	if err != nil {
		return nil, err
	}

	info.typeImporterTag = importer.tag
	return info, nil
}

func (inst *componentBuildingInfoMaker) parseTypeName(text string, info *componentBuildingInfo) error {

	text = strings.TrimSpace(text)
	index := strings.LastIndex(text, "/")

	if index < 0 {
		return errors.New("bad type name string:" + text)
	}

	str1 := text[0:index]
	str2 := text[index+1:]

	info.typeFullName = text
	info.typePackageName = str1
	info.typeShortName = str2

	return nil
}

func (inst *componentBuildingInfoMaker) parseAliases(text string) []string {
	return inst.parseStringArray(text)
}

func (inst *componentBuildingInfoMaker) parseClasses(text string) []string {
	return inst.parseStringArray(text)
}

func (inst *componentBuildingInfoMaker) parseStringArray(text string) []string {

	text = strings.ReplaceAll(text, "\t", ",")
	text = strings.ReplaceAll(text, " ", ",")
	array := strings.Split(text, ",")
	table := map[string]bool{}

	for index := range array {
		item := array[index]
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		table[item] = true
	}

	list := []string{}

	for key := range table {
		list = append(list, key)
	}

	return list
}
