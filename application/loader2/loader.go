package loader2

import (
	"errors"
	"os"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/vlog"
)

type loader struct {
}

func (inst *loader) _Impl() application.ContextLoader {
	return inst
}

func (inst *loader) Load(cfg application.Configuration, args []string) (application.Context, error) {
	loading := &contextLoading{}
	loading.init(cfg, args)
	return loading.load()
}

////////////////////////////////////////////////////////////////////////////////

type contextLoading struct {
	context   application.Context
	config    application.Configuration
	modules   []application.Module
	arguments []string
	profile   string // the value of ${application.profiles.active}
}

func (inst *contextLoading) init(cfg application.Configuration, args []string) {

	ctx := &appContext{}

	inst.modules = cfg.GetModules()
	inst.arguments = args
	inst.config = cfg
	inst.context = ctx.init()
	inst.context.SetErrorHandler(cfg.GetErrorHandler())
}

func (inst *contextLoading) load() (application.Context, error) {

	vlog.Debug("load args ...")
	err := inst.loadArguments()
	if err != nil {
		return nil, err
	}

	vlog.Debug("load env ...")
	err = inst.loadEnvironment()
	if err != nil {
		return nil, err
	}

	vlog.Debug("load resources ...")
	err = inst.loadResources()
	if err != nil {
		return nil, err
	}

	vlog.Debug("load properties ...")
	err = inst.loadProperties()
	if err != nil {
		return nil, err
	}

	vlog.Debug("load attributes ...")
	err = inst.loadAttributes()
	if err != nil {
		return nil, err
	}

	vlog.Debug("load logger ...")
	err = inst.loadLogger()
	if err != nil {
		return nil, err
	}

	vlog.Debug("load banner ...")
	err = inst.loadBanner()
	if err != nil {
		return nil, err
	}

	vlog.Debug("load about info ...")
	err = inst.loadAboutInfo()
	if err != nil {
		return nil, err
	}

	vlog.Debug("load components ...")
	err = inst.loadComponents()
	if err != nil {
		return nil, err
	}

	return inst.context, nil
}

func (inst *contextLoading) loadResources() error {
	src := inst.config.GetResources()
	dst := inst.context.GetResources()
	dst.Import(src.Export(nil), true)
	return nil
}

func (inst *contextLoading) loadLogger() error {
	loader := &loggerLoader{}
	return loader.load(inst)
}

func (inst *contextLoading) loadBanner() error {
	loader := &bannerLoader{}
	return loader.load(inst)
}

func (inst *contextLoading) loadAboutInfo() error {
	loader := &aboutInfoLoader{}
	return loader.load(inst)
}

func (inst *contextLoading) loadArguments() error {
	args := inst.arguments
	if args == nil {
		args = os.Args
	}
	ctx := inst.context
	ctx.GetArguments().Import(args)
	return nil
}

func (inst *contextLoading) loadEnvironment() error {
	ctx := inst.context
	dst := ctx.GetEnvironment()
	src := os.Environ()
	for _, kv := range src {
		index := strings.Index(kv, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(kv[0:index])
		val := strings.TrimSpace(kv[index+1:])
		dst.SetEnv(key, val)
	}
	return nil
}

func (inst *contextLoading) loadProperties() error {
	loader := &propertiesLoader{}
	p, err := loader.Load(inst)
	if err != nil {
		return err
	}
	dst := inst.context.GetProperties()
	dst.Import(p.Export(nil))
	inst.profile = loader.profile
	return nil
}

func (inst *contextLoading) loadAttributes() error {
	src := inst.config.GetAttributes()
	dst := inst.context.GetAttributes()
	if src == nil || dst == nil {
		return errors.New("atts==nil")
	}
	all := src.Export(nil)
	dst.Import(all)
	return nil
}

func (inst *contextLoading) loadComponents() error {
	subloader := &componentsLoader{}
	subloader.init(inst)
	return subloader.load()
}
