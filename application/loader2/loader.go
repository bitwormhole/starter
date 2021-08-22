package loader2

import (
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
	arguments []string
	profile   string // the value of ${application.profiles.active}
}

func (inst *contextLoading) init(cfg application.Configuration, args []string) {

	ctx := &appContext{}

	inst.arguments = args
	inst.config = cfg
	inst.context = ctx.init()
}

func (inst *contextLoading) load() (application.Context, error) {

	vlog.Debug("load properties ...")
	err := inst.loadProperties()
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

func (inst *contextLoading) loadProperties() error {
	loader := &propertiesLoader{}
	return loader.load(inst)
}

func (inst *contextLoading) loadComponents() error {
	subloader := &componentsLoader{}
	subloader.init(inst)
	return subloader.load()
}
