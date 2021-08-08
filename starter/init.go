package starter

import (
	"log"
	"os"
	"strconv"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/collection"
	etc "github.com/bitwormhole/starter/etc/starter"
)

// InitApp 开始初始化应用程序
func InitApp() application.Initializer {
	inst := &innerInitializer{}
	i := inst.init()
	i.Use(Module())
	return i
}

// Module 导出【starter】模块
func Module() application.Module {
	return etc.ExportModule()
}

////////////////////////////////////////////////////////////////////////////////

type innerInitializer struct {
	modules    *moduleManager
	cfgBuilder application.ConfigBuilder
}

// public

func (inst *innerInitializer) SetAttribute(name string, value interface{}) application.Initializer {
	inst.cfgBuilder.SetAttribute(name, value)
	return inst
}

func (inst *innerInitializer) Use(module application.Module) application.Initializer {
	inst.modules.use(module, true)
	return inst
}

func (inst *innerInitializer) Run() {
	err := inst.inTryRun()
	if err != nil {
		panic(err)
	}
}

// private

func (inst *innerInitializer) init() application.Initializer {
	inst.modules = createModuleManager()
	inst.cfgBuilder = config.NewBuilder()
	return inst
}

func (inst *innerInitializer) loadResourcesFromModules(mods []application.Module) error {
	sum := collection.CreateResources()
	for _, mod := range mods {
		res := mod.GetResources()
		if res == nil {
			continue
		}
		items := res.Export(nil)
		sum.Import(items, true)
	}
	inst.cfgBuilder.SetResources(sum)
	return nil
}

func (inst *innerInitializer) applyModules(mods []application.Module) error {
	cb := inst.cfgBuilder
	props := cb.DefaultProperties()
	for index, mod := range mods {
		log.Println("use module", mod.GetName(), mod.GetVersion())
		err := mod.Apply(cb)
		if err != nil {
			return err
		}
		inst.writeModuleInfoToProperties(props, index, mod)
	}
	return nil
}

func (inst *innerInitializer) writeModuleInfoToProperties(props collection.Properties, index int, mod application.Module) {

	name := mod.GetName()
	ver := mod.GetVersion()
	rev := strconv.Itoa(mod.GetRevision())
	idx := strconv.Itoa(index)

	prefix := "module." + idx + "."

	props.SetProperty(prefix+"name", name)
	props.SetProperty(prefix+"version", ver)
	props.SetProperty(prefix+"revision", rev)
	props.SetProperty(prefix+"index", idx)
}

// func (inst *innerInitializer) useDependencies(deps []application.Module) {
// 	if deps == nil {
// 		return
// 	}
// 	for index := range deps {
// 		item := deps[index]
// 		inst.Use(item)
// 	}
// }

func (inst *innerInitializer) inTryRun() error {

	mods := inst.modules.listAll()

	err := inst.loadResourcesFromModules(mods)
	if err != nil {
		return err
	}

	err = inst.applyModules(mods)
	if err != nil {
		return err
	}

	cfg := inst.cfgBuilder.Create()
	context, err := application.Run(cfg, os.Args)
	if err != nil {
		return err
	}

	err = application.Loop(context)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	if err != nil {
		return err
	}

	log.Println("exit with code:", code)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
