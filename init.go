package starter

import (
	"os"
	"runtime"
	"runtime/debug"
	"strconv"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/bootstrap"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/starter/vlog/std"
)

// InitApp 开始初始化应用程序
func InitApp() application.Initializer {
	inst := &innerInitializer{}
	i := inst.init()
	i.Use(Module())
	i.UsePanic()
	i.SetExitEnabled(true)
	return i
}

////////////////////////////////////////////////////////////////////////////////

type innerInitializer struct {
	modules       *moduleManager
	cfgBuilder    application.ConfigBuilder
	errHandler    lang.ErrorHandler
	osExitEnabled bool
}

// public

func (inst *innerInitializer) SetAttribute(name string, value interface{}) application.Initializer {
	inst.cfgBuilder.SetAttribute(name, value)
	return inst
}

func (inst *innerInitializer) SetErrorHandler(h lang.ErrorHandler) application.Initializer {
	inst.cfgBuilder.SetErrorHandler(h)
	inst.errHandler = h
	return inst
}

func (inst *innerInitializer) SetExitEnabled(en bool) application.Initializer {
	inst.osExitEnabled = en
	return inst
}

func (inst *innerInitializer) SetPanicEnabled(en bool) application.Initializer {
	if en {
		inst.UsePanic()
	} else {
		inst.SetErrorHandler(nil)
	}
	return inst
}

func (inst *innerInitializer) Use(module application.Module) application.Initializer {
	inst.modules.use(module, true)
	return inst
}

func (inst *innerInitializer) UseProperties(p collection.Properties) application.Initializer {
	inst.cfgBuilder.AddProperties(p)
	return inst
}

func (inst *innerInitializer) UseResources(r collection.Resources) application.Initializer {
	inst.cfgBuilder.AddResources(r)
	return inst
}

func (inst *innerInitializer) UsePanic() application.Initializer {
	h := lang.NewErrorHandlerForFunc(func(err error) error {
		panic(err) // return err
	})
	inst.SetErrorHandler(h)
	return inst
}

func (inst *innerInitializer) Run() {
	err := inst.inTryRun()
	if err != nil {
		panic(err)
	}
}

func (inst *innerInitializer) RunEx() (application.Runtime, error) {
	rt1 := &innerAppRuntime{}
	rt2, err := rt1.initAndRun(inst)
	if err != nil {
		return nil, err
	}
	return rt2, nil
}

// private

func (inst *innerInitializer) init() application.Initializer {
	inst.initLogging()
	inst.modules = createModuleManager()
	inst.cfgBuilder = bootstrap.ConfigBuilder() // config.NewBuilder()
	inst.loadBasicProperties()
	return inst
}

func (inst *innerInitializer) initLogging() {
	vlog.SetDefaultFactory(&std.StandardLoggerFactory{})
}

func (inst *innerInitializer) loadBasicProperties() {

	appname := inst.loadPropAppName()
	goVer := inst.loadPropGoVersion()
	hostname := inst.loadPropHostName()

	dp := inst.cfgBuilder.DefaultProperties()
	dp.SetProperty("go.version", goVer)
	dp.SetProperty("application.name", appname)
	dp.SetProperty("host.name", hostname)

}

func (inst *innerInitializer) loadPropGoVersion() string {
	return runtime.Version()
}

func (inst *innerInitializer) loadPropAppName() string {
	info, ok := debug.ReadBuildInfo()
	if ok {
		return info.Main.Path
	}
	return "unnamed"
}

func (inst *innerInitializer) loadPropHostName() string {
	name, err := os.Hostname()
	if err == nil {
		return name
	}
	return "localhost"
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
		vlog.Info("use module ", mod.GetName(), "@", mod.GetVersion())
		err := mod.Apply(cb)
		if err != nil {
			return err
		}
		inst.writeModuleInfoToProperties(props, index, mod)
		inst.tryLoadDefaultPropertiesFromRes(props, index, mod)
	}
	return nil
}

func (inst *innerInitializer) tryLoadDefaultPropertiesFromRes(props collection.Properties, index int, mod application.Module) {

	vlog.Trace("try load 'default.properties' from module: " + mod.GetName())

	r := mod.GetResources()
	if r == nil {
		return
	}

	text, err := r.GetText("default.properties")
	if err != nil {
		vlog.Warn(err, ", mod=", mod.GetName())
		return
	}

	p, err := collection.ParseProperties(text, nil)
	if err != nil {
		vlog.Warn(err, ", mod=", mod.GetName())
		return
	}

	props.Import(p.Export(nil))
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

func (inst *innerInitializer) inTryRun() error {

	rt, err := inst.RunEx()
	if err != nil {
		return err
	}

	err = rt.Loop()
	if err != nil {
		return err
	}

	err = rt.Exit()
	if err != nil {
		return err
	}

	return nil
}

func (inst *innerInitializer) prepareConfig() (application.Configuration, error) {

	mods := inst.modules.listAll()

	err := inst.loadResourcesFromModules(mods)
	if err != nil {
		return nil, err
	}

	err = inst.applyModules(mods)
	if err != nil {
		return nil, err
	}

	cfg := inst.cfgBuilder.Create()
	return cfg, nil
}

func (inst *innerInitializer) handleError(err error) error {
	if err == nil {
		return nil
	}
	h := inst.errHandler
	if h == nil {
		return err
	}
	return h.HandleError(err)
}

////////////////////////////////////////////////////////////////////////////////

type innerAppRuntime struct {
	parent  *innerInitializer
	context application.Context
}

func (inst *innerAppRuntime) initAndRun(parent *innerInitializer) (application.Runtime, error) {

	cfg, err := parent.prepareConfig()
	if err != nil {
		err = parent.handleError(err)
		return nil, err
	}

	ctx, err := application.Run(cfg, os.Args)
	if err != nil {
		err = parent.handleError(err)
		return nil, err
	}

	inst.context = ctx
	inst.parent = parent
	return inst, nil
}

func (inst *innerAppRuntime) Context() application.Context {
	return inst.context
}

func (inst *innerAppRuntime) Loop() error {
	ctx := inst.context
	err := application.Loop(ctx)
	if err != nil {
		err = inst.parent.handleError(err)
		return err
	}
	return nil
}

func (inst *innerAppRuntime) Exit() error {
	ctx := inst.context
	code, err := application.Exit(ctx)
	if err != nil {
		err = inst.parent.handleError(err)
		return err
	}
	vlog.Info("exit with code:", code)
	if inst.parent.osExitEnabled {
		os.Exit(code)
	}
	return nil
}
