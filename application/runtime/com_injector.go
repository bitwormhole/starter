package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////

type innerInjector struct {
}

func (inst *innerInjector) init() application.Injector {
	return inst
}

func (inst *innerInjector) OpenInjection(ctx application.Context) (application.Injection, error) {
	comLoader := ctx.ComponentLoader()
	comLoading, err := comLoader.OpenLoading(ctx)
	if err != nil {
		return nil, err
	}
	injection := &innerInjection{}
	injection.init(comLoading)
	return injection, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerInjection struct {
	pool    lang.ReleasePool
	context application.Context
	loading application.ComponentLoading
}

func (inst *innerInjection) init(loading application.ComponentLoading) application.Injection {
	inst.pool = loading.Pool()
	inst.context = loading.Context()
	return inst
}

func (inst *innerInjection) OnError(err error) {
	inst.loading.OnError(err)
}

func (inst *innerInjection) Pool() lang.ReleasePool {
	return inst.pool
}

func (inst *innerInjection) Context() application.Context {
	return inst.context
}

func (inst *innerInjection) Select(selector string) application.InjectionSource {
	return nil
}

func (inst *innerInjection) Close() error {
	return inst.loading.Close()
}
