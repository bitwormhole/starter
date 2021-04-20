package configen

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

type singleConfigFile struct {
	context         *configenContext
	path            fs.Path
	name            string
	text            string
	keyPrefix       string
	properties      collection.Properties
	includeFileList []string
	head            *configHead
	loaded          bool
}

func (inst *singleConfigFile) load() error {

	if inst.loaded {
		return nil
	} else {
		inst.loaded = true
	}

	fmt.Println("load component config from " + inst.path.Path())

	text, err := inst.path.GetIO().ReadText()
	if err != nil {
		return err
	}

	text = inst.setupKeyForComponents(text)
	inst.text = text
	inst.properties, err = collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	inst.head, err = inst.loadHead()
	if err != nil {
		return err
	}

	inst.includeFileList, err = inst.loadIncludeFileList()
	if err != nil {
		return err
	}

	return nil
}

func (inst *singleConfigFile) loadHead() (*configHead, error) {

	props := inst.properties
	head := &configHead{}

	head.configFunctionName = props.GetProperty("head.configFunction", "")
	head.packageName = props.GetProperty("head.package", "")

	if head.packageName == "" {
		return nil, errors.New("no config value: [head.package]")
	}

	if head.configFunctionName == "" {
		return nil, errors.New("no config value: [head.configFunction]")
	}

	return head, nil
}

func (inst *singleConfigFile) loadIncludeFileList() ([]string, error) {
	value := inst.properties.GetProperty("head.include", "")
	list1 := strings.Split(value, ",")
	list2 := []string{}
	for index := range list1 {
		item := strings.TrimSpace(list1[index])
		if item == "" {
			continue
		}
		list2 = append(list2, item)
	}
	return list2, nil
}

func (inst *singleConfigFile) setupKeyForComponents(text string) string {

	array := strings.Split(text, "[component]")
	builder := &strings.Builder{}

	for index := range array {
		if index > 0 {
			builder.WriteString("[component \"")
			builder.WriteString(inst.keyPrefix)
			builder.WriteString("com")
			builder.WriteString(strconv.Itoa(index))
			builder.WriteString("\"]")
		}
		builder.WriteString(array[index])
	}

	return builder.String()
}
