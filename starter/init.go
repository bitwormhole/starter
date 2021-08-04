package starter

import (
	"embed"
	"log"
	"os"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/collection"
	etc "github.com/bitwormhole/starter/etc/starter"
)

// InitApp 开始初始化应用程序
func InitApp() application.Initializer {
	inst := &innerInitializer{}
	i := inst.init()
	i.Use(etc.Module())
	return i
}

////////////////////////////////////////////////////////////////////////////////

type innerInitializer struct {
	modules    map[string]application.Module
	cfgBuilder application.ConfigBuilder
}

// public

func (inst *innerInitializer) EmbedResources(fs *embed.FS, path string) application.Initializer {
	res := config.CreateEmbedFsResources(fs, path)
	inst.cfgBuilder.SetResources(res)
	return inst
}

func (inst *innerInitializer) MountResources(res collection.Resources, path string) application.Initializer {
	inst.cfgBuilder.SetResources(res)
	return inst
}

func (inst *innerInitializer) SetAttribute(name string, value interface{}) application.Initializer {
	inst.cfgBuilder.SetAttribute(name, value)
	return inst
}

func (inst *innerInitializer) Use(module application.Module) application.Initializer {
	if module == nil {
		return inst
	}
	name := module.GetName()
	older := inst.modules[name]
	if older != nil {
		if older.GetRevision() >= module.GetRevision() {
			return inst
		}
	}
	inst.modules[name] = module
	inst.useDependencies(module.GetDependencies())
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
	inst.modules = make(map[string]application.Module)
	inst.cfgBuilder = config.NewBuilder()
	return inst
}

func (inst *innerInitializer) applyModules() error {
	mods := inst.modules
	cb := inst.cfgBuilder
	for key := range mods {
		mod := mods[key]
		log.Println("use module", mod.GetName(), mod.GetVersion())
		err := mod.Apply(cb)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *innerInitializer) useDependencies(deps []application.Module) {
	if deps == nil {
		return
	}
	for index := range deps {
		item := deps[index]
		inst.Use(item)
	}
}

func (inst *innerInitializer) inTryRun() error {

	err := inst.applyModules()
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
