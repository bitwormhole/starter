package configen

import (
	"errors"
	"strconv"
	"strings"
)

type importerBuildingInfo struct {
	fullName string
	tag      string // default: the suffix of fullName
}

type importerManager struct {
	context *configenContext
	tags    map[string]*importerBuildingInfo
}

func (inst *importerManager) init() {
	if inst.tags == nil {
		inst.tags = map[string]*importerBuildingInfo{}
	}
}

func (inst *importerManager) loadImporter(pkgName string) (*importerBuildingInfo, error) {
	item := inst.context.importerTable[pkgName]
	if item != nil {
		return item, nil
	}
	tag := inst.createTagForPackageName(pkgName)
	item = &importerBuildingInfo{}
	item.fullName = pkgName
	item.tag = tag
	// save
	inst.tags[tag] = item
	inst.context.importerTable[pkgName] = item
	return item, nil
}

func (inst *importerManager) getImporter(pkgName string) (*importerBuildingInfo, error) {
	item := inst.context.importerTable[pkgName]
	if item == nil {
		return nil, errors.New("no import-info with name: " + pkgName)
	}
	return item, nil
}

func (inst *importerManager) createTagForPackageName(packname string) string {

	inst.init()
	index := strings.LastIndex(packname, "/")
	tag := ""

	if index < 0 {
		tag = packname
	} else {
		tag = packname[index+1:]
	}

	tag = strings.ReplaceAll(tag, "-", "_")
	tag = strings.ReplaceAll(tag, ".", "_")
	tag0 := tag

	for index = 1; ; index++ {
		importer := inst.tags[tag]
		if importer == nil {
			break
		}
		tag = tag0 + strconv.Itoa(index)
	}

	return tag
}
