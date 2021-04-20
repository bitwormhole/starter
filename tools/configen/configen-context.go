package configen

import (
	"strconv"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

type configenContext struct {
	inputFileName  string
	outputFileName string
	pwd            fs.Path
	inputFile      fs.Path
	outputFile     fs.Path
	head           *configHead
	properties     collection.Properties // 把所有的 singleConfigFile 混合于此
	code           string

	components          []*componentDescriptor
	comBuildingInfoList []*componentBuildingInfo
	singleConfigFiles   map[string]*singleConfigFile
	importerTable       map[string]*importerBuildingInfo

	importerManager *importerManager
}

func (inst *configenContext) getSingleConfigFile(name string, create bool) *singleConfigFile {
	table := inst.singleConfigFiles
	item := table[name]
	if item == nil {
		if create {
			size := len(table)
			item = &singleConfigFile{
				context: inst,
				name:    name,
			}
			item.keyPrefix = "file" + strconv.Itoa(size)
			item.path = inst.pwd.GetChild(name)
			table[name] = item
		}
	}
	return item
}
