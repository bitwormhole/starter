package configenchecker

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

type ConfigenChecker struct {
	Context application.Context
	Enable  bool
}

func (inst *ConfigenChecker) Check() error {
	if inst.Enable {
		return inst.doCheck()
	}
	return nil
}

func (inst *ConfigenChecker) doCheck() error {

	exe := os.Args[0]
	path0 := fs.Default().GetPath(exe)
	log.Println("CHECK configen.properties in ", path0.Path())

	dir, err := inst.findProjectDir(path0)
	if err != nil {
		// return err
		return nil // 找不到就算了，反正这个检查不是强制的:-(
	}

	nodes, err := inst.loadConfigenNodesInProjectDir(dir)
	if err != nil {
		return err
	}

	return inst.checkConfigenNodes(nodes)
}

// 根据给定的路径，查找 go.mod 项目文件夹
func (inst *ConfigenChecker) findProjectDir(from fs.Path) (fs.Path, error) {
	path := from
	const targetFile = "go.mod"
	for tout := 9; tout > 0; tout-- {
		if path == nil {
			break
		}
		child := path.GetChild(targetFile)
		if child.Exists() && child.IsFile() {
			return path, nil
		}
		path = path.Parent()
	}
	msg := "cannot find [" + targetFile + "] in path [" + from.Path() + "]"
	return nil, errors.New(msg)
}

// 从项目文件夹里面的根“configen.properties”文件，加载该项目中的所有的“configen.properties”节点
func (inst *ConfigenChecker) loadConfigenNodesInProjectDir(dir fs.Path) ([]fs.Path, error) {

	const filename = "configen.properties"
	const keyPrefix = "node."
	const keySuffix = ".path"

	// 加载properties
	file := dir.GetChild(filename)
	props, err := inst.loadPropertiesInFile(file)
	if err != nil {
		return nil, err
	}

	// 遍历properties
	kvTab := props.Export(nil)
	results := make([]fs.Path, 0)
	for key := range kvTab {
		if strings.HasPrefix(key, keyPrefix) && strings.HasSuffix(key, keySuffix) {
			value := kvTab[key]
			child := dir.GetChild(value)
			results = append(results, child)
		}
	}

	return results, nil
}

// 检查所有的节点
func (inst *ConfigenChecker) checkConfigenNodes(nodes []fs.Path) error {
	for index := range nodes {
		node := nodes[index]
		err := inst.checkConfigenNode(node)
		if err != nil {
			return err
		}
	}
	log.Println("[OK]")
	return nil
}

func (inst *ConfigenChecker) checkConfigenNode(file fs.Path) error {

	path := file.Path()
	dir := file.Parent()
	const keyOutputFile = "configen.output.file"
	log.Println("  check ", path)

	props, err := inst.loadPropertiesInFile(file)
	if err != nil {
		return err
	}

	outputFileName, err := props.GetPropertyRequired(keyOutputFile)
	if err != nil {
		return err
	}

	outputFile := dir.GetChild(outputFileName)
	outputFileModTime := outputFile.LastModTime()

	// 检查自动生成的文件（auto_generated_by_starter_configen.go）是否是最新的

	items := dir.ListItems()
	for index := range items {
		item := items[index]
		modTime := item.LastModTime()
		if modTime <= outputFileModTime {
			continue
		} else {
			msg := "the file is updated, need for starter-configen rebuild. file=" + item.Path()
			return errors.New(msg)
		}
	}

	return nil
}

func (inst *ConfigenChecker) loadPropertiesInFile(file fs.Path) (collection.Properties, error) {
	text, err := file.GetIO().ReadText()
	if err != nil {
		return nil, err
	}
	return collection.ParseProperties(text, nil)
}
