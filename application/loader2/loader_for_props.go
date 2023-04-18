package loader2

import (
	"os"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

const (
	ProfileNameDefault = "[default]"
	ProfileNameApp     = "[app]"
)

////////////////////////////////////////////////////////////////////////////////

// 属性加载器
type propertiesLoader struct {

	// cache
	cacheForArgs   collection.Properties
	cacheForExeDir collection.Properties
	cacheForMods   map[string]collection.Properties

	profile string
	mods    []application.Module
	args    []string
	plist   []collection.Properties
}

func (inst *propertiesLoader) Load(loading *contextLoading) (collection.Properties, error) {

	inst.mods = loading.modules
	inst.args = loading.arguments
	inst.plist = nil
	inst.cacheForMods = make(map[string]collection.Properties)
	inst.profile = ""

	err := inst.applySteps()
	if err != nil {
		return nil, err
	}

	inst.profile = inst.computeProfileName()
	inst.plist = nil

	err = inst.applySteps()
	if err != nil {
		return nil, err
	}

	return inst.makeResult()
}

func (inst *propertiesLoader) makeResult() (collection.Properties, error) {
	list := inst.plist
	dst := collection.CreateProperties()
	for _, src := range list {
		table := src.Export(nil)
		dst.Import(table)
	}
	return dst, nil
}

func (inst *propertiesLoader) applySteps() error {

	steps := make([]func() error, 0)

	steps = append(steps, inst.loadFromModules)
	steps = append(steps, inst.loadFromExeDir)
	steps = append(steps, inst.loadFromArgs)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *propertiesLoader) computeProfileName() string {
	const key = "application.profiles.active"
	list := inst.plist
	value := ""
	for _, p := range list {
		value = p.GetProperty(key, value)
	}
	return value
}

func (inst *propertiesLoader) loadFromArgs() error {

	// by cache
	props := inst.cacheForArgs
	if props != nil {
		inst.plist = append(inst.plist, props)
		return nil
	}

	// do load
	const prefix = "--"
	args := inst.args
	props = collection.CreateProperties()

	for _, item := range args {
		if !strings.HasPrefix(item, prefix) {
			continue
		}
		item = item[len(prefix):]
		index := strings.Index(item, "=")
		if index < 1 {
			continue
		}
		key := strings.TrimSpace(item[0:index])
		val := strings.TrimSpace(item[index+1:])
		props.SetProperty(key, val)
	}

	inst.cacheForArgs = props
	inst.plist = append(inst.plist, props)
	return nil
}

func (inst *propertiesLoader) loadFromExeDir() error {

	// by cache
	props := inst.cacheForExeDir
	if props != nil {
		inst.plist = append(inst.plist, props)
		return nil
	}

	// do load
	exepath, err := os.Executable()
	if err != nil {
		return err
	}

	exefile := fs.Default().GetPath(exepath)
	pfile := exefile.Parent().GetChild("application.properties")
	if !pfile.IsFile() {
		return nil
	}

	text, err := pfile.GetIO().ReadText(nil)
	props = collection.CreateProperties()
	props, err = collection.ParseProperties(text, props)
	if err != nil {
		return err
	}

	inst.cacheForExeDir = props
	inst.plist = append(inst.plist, props)
	return nil
}

func (inst *propertiesLoader) loadFromModules() error {
	profile := inst.profile
	mods := inst.mods
	for _, mod := range mods {
		inst.tryLoadFromRes(ProfileNameDefault, mod)
		inst.tryLoadFromRes(ProfileNameApp, mod)
		inst.tryLoadFromRes(profile, mod)
	}
	return nil
}

func (inst *propertiesLoader) tryLoadFromRes(profile string, mod application.Module) {
	err := inst.doLoadFromRes(profile, mod)
	if err != nil {
		vlog.Warn(err)
	}
}

func (inst *propertiesLoader) doLoadFromRes(profile string, mod application.Module) error {

	url := ""
	if profile == "" {
		return nil // skip

	} else if profile == ProfileNameApp {
		url = "res:///application.properties"

	} else if profile == ProfileNameDefault {
		url = "res:///default.properties"

	} else {
		url = "res:///application-" + profile + ".properties"
	}

	// by cache
	nameForCache := url + "?module=" + mod.GetName()
	props := inst.cacheForMods[nameForCache]
	if props != nil {
		inst.plist = append(inst.plist, props)
		return nil
	}

	// do load
	res := mod.GetResources()
	if res == nil {
		return nil // skip
	}

	text, err := res.GetText(url)
	if err != nil {
		return err
	}

	props, err = collection.ParseProperties(text, nil)
	if err != nil {
		return nil // err
	}

	// save
	inst.cacheForMods[nameForCache] = props
	inst.plist = append(inst.plist, props)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
