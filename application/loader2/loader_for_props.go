package loader2

import (
	"errors"
	"os"
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

// 属性加载器
type propertiesLoader struct {
	loading *contextLoading
}

func (inst *propertiesLoader) loadFromArgs() error {
	ctx := inst.loading.context
	args := ctx.GetArguments()
	props := ctx.GetProperties()
	list := args.Export()
	for _, item := range list {
		if !strings.HasPrefix(item, "--") {
			continue
		}
		item = item[2:]
		index := strings.Index(item, "=")
		if index < 1 {
			continue
		}
		key := strings.TrimSpace(item[0:index])
		val := strings.TrimSpace(item[index+1:])
		props.SetProperty(key, val)
	}
	return nil
}

func (inst *propertiesLoader) loadFromRes(name string) error {
	ctx := inst.loading.context
	res := ctx.GetResources()
	props := ctx.GetProperties()
	text, err := res.GetText("/" + name)
	if err != nil {
		return nil
	}
	collection.ParseProperties(text, props)
	inst.loadFromArgs()
	return nil
}

func (inst *propertiesLoader) loadFromRes1() error {
	return inst.loadFromRes("application.properties")
}

func (inst *propertiesLoader) loadFromRes2() error {
	const key = "application.profiles.active"
	ctx := inst.loading.context
	profile := ctx.GetProperties().GetProperty(key, "default")
	inst.loading.profile = profile
	return inst.loadFromRes("application-" + profile + ".properties")
}

// 取指定的属性，并转为 bool 值
func (inst *propertiesLoader) getBool(name string, p collection.Properties) bool {
	if p == nil {
		return false
	}
	text, err := p.GetPropertyRequired(name)
	if err != nil {
		return false
	}
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	if text == "true" || text == "1" || text == "yes" {
		return true
	}
	return false
}

// 从指定的文件加载属性，存入到dst
func (inst *propertiesLoader) loadFromFilePath(path string, dst collection.Properties) error {

	file := fs.Default().GetPath(path)

	if !file.IsFile() {
		return errors.New("the file is not exists, path=" + file.Path())
	}

	vlog.Info("load application.Properties from ", file.Path())
	text, err := file.GetIO().ReadText(nil)

	if err != nil {
		return nil
	}

	_, err = collection.ParseProperties(text, dst)
	return err
}

func (inst *propertiesLoader) loadFromFile() error {

	const keyPath = "application.properties.file.path"
	const keyEnabled = "application.properties.file.enabled"
	const keyRequired = "application.properties.file.required"

	ctx := inst.loading.context
	props := ctx.GetProperties()

	enabled := inst.getBool(keyEnabled, props)
	required := inst.getBool(keyRequired, props)

	if !enabled {
		return nil
	}

	path, err := props.GetPropertyRequired(keyPath)
	if err != nil {
		if required {
			return err
		}
		return nil
	}

	err = inst.loadFromFilePath(path, props)
	if err != nil {
		if required {
			return err
		}
		return nil
	}

	inst.loadFromArgs()
	return nil
}

func (inst *propertiesLoader) loadFromDefault() error {
	cfg := inst.loading.config
	ctx := inst.loading.context
	src := cfg.GetDefaultProperties()
	dst := ctx.GetProperties()
	dst.Import(src.Export(nil))
	return nil
}

func (inst *propertiesLoader) loadFromExeDir() error {

	exe, err := os.Executable()
	if err != nil {
		return err
	}

	exefile := fs.Default().GetPath(exe)
	pfile := exefile.Parent().GetChild("application.properties")
	if !pfile.IsFile() {
		return nil
	}

	ctx := inst.loading.context
	props := ctx.GetProperties()
	return inst.loadFromFilePath(pfile.Path(), props)
}

func (inst *propertiesLoader) loadFromFinal() error {
	cfg := inst.loading.config
	ctx := inst.loading.context
	src := cfg.GetFinalProperties()
	dst := ctx.GetProperties()
	dst.Import(src.Export(nil))
	return nil
}

func (inst *propertiesLoader) load(loading *contextLoading) error {

	inst.loading = loading

	err := inst.loadFromDefault()
	if err != nil {
		return err
	}

	err = inst.loadFromArgs()
	if err != nil {
		return err
	}

	err = inst.loadFromRes1()
	if err != nil {
		return err
	}

	err = inst.loadFromRes2()
	if err != nil {
		return err
	}

	err = inst.loadFromFile()
	if err != nil {
		return err
	}

	err = inst.loadFromExeDir()
	if err != nil {
		return err
	}

	err = inst.loadFromFinal()
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
