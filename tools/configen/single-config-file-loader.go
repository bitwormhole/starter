package configen

import "errors"

type singleConfigFileLoader struct {
	context *configenContext
}

func (inst *singleConfigFileLoader) loadAllSingleFiles() error {
	name := inst.context.inputFileName
	item := inst.context.getSingleConfigFile(name, true)
	err := inst.loadSingleFile(item, 10)
	if err != nil {
		return err
	}
	inst.context.head = item.head
	return nil
}

func (inst *singleConfigFileLoader) loadSingleFile(target *singleConfigFile, depth_limit int) error {

	if depth_limit < 0 {
		return errors.New("the include recursion is too deep.")
	}

	if target.loaded {
		return nil
	}

	err := target.load()
	if err != nil {
		return err
	}

	inst.putItemsToMaster(target)
	include_list := target.includeFileList

	for index := range include_list {
		name := include_list[index]
		item := inst.context.getSingleConfigFile(name, true)
		err = inst.loadSingleFile(item, depth_limit-1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *singleConfigFileLoader) putItemsToMaster(target *singleConfigFile) {
	src := target.properties.Export(nil)
	dst := inst.context.properties
	for key := range src {
		value := src[key]
		dst.SetProperty(key, value)
	}
}
