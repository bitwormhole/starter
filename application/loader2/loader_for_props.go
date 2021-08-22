package loader2

import (
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

func (inst *propertiesLoader) loadFromFile() error {

	const key = "application.properties.file"
	ctx := inst.loading.context
	props := ctx.GetProperties()

	path, err := props.GetPropertyRequired(key)
	if err != nil {
		return nil
	}

	file := fs.Default().GetPath(path)
	if !file.IsFile() {
		return nil
	}
	vlog.Info("load application.Properties from ", file.Path())

	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil
	}
	collection.ParseProperties(text, props)

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

	return nil
}

////////////////////////////////////////////////////////////////////////////////
