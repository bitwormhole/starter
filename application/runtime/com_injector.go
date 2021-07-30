package runtime

import (
	"errors"
	"strconv"
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

func (inst *innerInjection) loadAllComponents(holders []application.ComponentHolder, source *innerInjectionSource) {
	comlist, err := inst.loading.LoadAll(holders)
	if err != nil {
		inst.loading.OnError(err)
		return
	}
	for index := range comlist {
		obj := comlist[index]
		source.add(obj)
	}
}

func (inst *innerInjection) selectByID(selector string, source *innerInjectionSource) {
	source.init(selector, "[id]")
	holder, err := inst.context.GetComponents().GetComponent(selector)
	inst.loading.OnError(err)
	if err == nil {
		com, err := inst.loading.Load(holder)
		inst.loading.OnError(err)
		if err == nil {
			source.add(com)
		}
	}
}

func (inst *innerInjection) selectAll(selector string, source *innerInjectionSource) {
	source.init(selector, "[*]")
	allcom := inst.context.GetComponents()
	holders := allcom.GetComponentsByFilter(func(name string, holder application.ComponentHolder) bool {
		info := holder.GetInfo()
		if name != info.GetID() {
			return false
		}
		return true
	})
	inst.loadAllComponents(holders, source)
}

func (inst *innerInjection) selectByClass(selector string, source *innerInjectionSource) {
	source.init(selector, "[class]")
	holders := inst.context.GetComponents().GetComponents(selector)
	inst.loadAllComponents(holders, source)
}

func (inst *innerInjection) selectByScope(selector string, source *innerInjectionSource, scope application.ComponentScope) {
	source.init(selector, "[*.scope]")
	allcom := inst.context.GetComponents()
	holders := allcom.GetComponentsByFilter(func(name string, holder application.ComponentHolder) bool {
		info := holder.GetInfo()
		if name != info.GetID() {
			return false
		}
		scope2 := info.GetScope()
		return scope2 == scope
	})
	inst.loadAllComponents(holders, source)
}

func (inst *innerInjection) selectByProperty(selector string, source *innerInjectionSource) {
	i1 := strings.IndexByte(selector, '{')
	i2 := strings.LastIndexByte(selector, '}')
	if (0 < i1) && (i1 < i2) {
		key := selector[i1+1 : i2]
		value, err := inst.context.GetProperties().GetPropertyRequired(key)
		if err == nil {
			source.init(selector, value)
		} else {
			inst.loading.OnError(err)
		}
	} else {
		err := errors.New("bad property selector:" + selector)
		inst.loading.OnError(err)
	}
}

func (inst *innerInjection) Select(selector string) application.InjectionSource {

	source := &innerInjectionSource{selector: selector}

	if selector == "*" {
		// All
		inst.selectAll(selector, source)

	} else if strings.HasPrefix(selector, "#") {
		// ID
		inst.selectByID(selector, source)

	} else if strings.HasPrefix(selector, ".") {
		// Class
		inst.selectByClass(selector, source)

	} else if strings.HasPrefix(selector, "${") && strings.HasSuffix(selector, "}") {
		// Property
		inst.selectByProperty(selector, source)

	} else if selector == "*.scope(singleton)" {
		// *.Scope
		inst.selectByScope(selector, source, application.ScopeSingleton)

	} else {
		source.init(selector, "[unsupported]")
		err := errors.New("unsupported selector:" + selector)
		inst.loading.OnError(err)
	}

	source.prepare()
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
	text     string
}

func (inst *innerInjectionSource) init(selector string, value string) application.InjectionSource {
	inst.items = make([]lang.Object, 0)
	inst.selector = selector
	inst.ptr = 0
	inst.count = 0
	inst.text = value
	return inst
}

func (inst *innerInjectionSource) add(o lang.Object) {
	if o == nil {
		return
	}
	inst.items = append(inst.items, o)
}

func (inst *innerInjectionSource) setText(text string) {
	inst.text = text
}

func (inst *innerInjectionSource) prepare() {
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

func (inst *innerInjectionSource) ReadInt() (int, error) {
	n, err := strconv.ParseInt(inst.text, 0, 0)
	err = inst.formatError(err)
	return int(n), err
}

func (inst *innerInjectionSource) ReadInt32() (int32, error) {
	n, err := strconv.ParseInt(inst.text, 0, 32)
	err = inst.formatError(err)
	return int32(n), nil
}

func (inst *innerInjectionSource) ReadInt64() (int64, error) {
	n, err := strconv.ParseInt(inst.text, 0, 64)
	err = inst.formatError(err)
	return n, err
}

func (inst *innerInjectionSource) ReadFloat32() (float32, error) {
	n, err := strconv.ParseFloat(inst.text, 32)
	err = inst.formatError(err)
	return float32(n), err
}

func (inst *innerInjectionSource) ReadFloat64() (float64, error) {
	n, err := strconv.ParseFloat(inst.text, 64)
	err = inst.formatError(err)
	return n, err
}

func (inst *innerInjectionSource) ReadBool() (bool, error) {
	value, err := strconv.ParseBool(inst.text)
	err = inst.formatError(err)
	return value, err
}

func (inst *innerInjectionSource) formatError(err error) error {
	if err == nil {
		return nil
	}
	text1 := err.Error()
	text2 := inst.selector
	return errors.New(text1 + ", selector=" + text2)
}

func (inst *innerInjectionSource) ReadString() (string, error) {
	return inst.text, nil
}

func (inst *innerInjectionSource) Close() error {
	inst.text = ""
	// inst.selector = ""
	inst.items = []lang.Object{}
	inst.count = 0
	inst.ptr = 0
	return nil
}
