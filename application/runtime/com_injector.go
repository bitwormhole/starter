package runtime

import (
	"errors"
	"strings"

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
	inst.loading = loading
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

	aslist := false
	source := &innerInjectionSource{}
	source.init(selector)

	if strings.HasPrefix(selector, "#") {
		aslist = false
	} else {
		aslist = true
	}

	if aslist {
		holder, err := inst.context.GetComponents().GetComponent(selector)
		inst.loading.OnError(err)
		if err == nil {
			com, err := inst.loading.Load(holder)
			inst.loading.OnError(err)
			if err == nil {
				source.add(com)
			}
		}
	} else {
		holders := inst.context.GetComponents().GetComponents(selector)
		comlist, err := inst.loading.LoadAll(holders)
		inst.loading.OnError(err)
		for index := range comlist {
			com := comlist[index]
			source.add(com)
		}
	}

	source.reset()
	return source
}

func (inst *innerInjection) Close() error {
	return inst.loading.Close()
}

////////////////////////////////////////////////////////////////////////////////
// innerInjectionSource

type innerInjectionSource struct {
	items    []lang.Object
	count    int
	ptr      int
	selector string
}

func (inst *innerInjectionSource) init(selector string) application.InjectionSource {
	inst.items = make([]lang.Object, 0)
	inst.selector = selector
	inst.ptr = 0
	inst.count = 0
	return inst
}

func (inst *innerInjectionSource) add(o lang.Object) {
	if o == nil {
		return
	}
	inst.items = append(inst.items, o)
}

func (inst *innerInjectionSource) reset() {
	inst.ptr = 0
	inst.count = len(inst.items)
}

func (inst *innerInjectionSource) Count() int {
	return inst.count
}

func (inst *innerInjectionSource) HasMore() bool {
	cnt := inst.count
	ptr := inst.ptr
	return (0 <= ptr) && (ptr < cnt)
}

func (inst *innerInjectionSource) Selector() string {
	return inst.selector
}

func (inst *innerInjectionSource) Read() (lang.Object, error) {
	list := inst.items
	size := inst.count
	ptr := inst.ptr
	if list == nil {
		return nil, errors.New("com.list==nil")
	}
	if (0 <= ptr) && (ptr < size) {
		com := list[ptr]
		if com == nil {
			return nil, errors.New("com.item==nil")
		}
		inst.ptr = ptr + 1
		return com, nil
	} else {
		return nil, errors.New("ptr out of bounds.")
	}
}

func (inst *innerInjectionSource) Close() error {
	inst.init("[closed]")
	return nil
}
