package loader2

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

type comInstanceHolder struct {
	id       string
	holder   application.ComponentHolder
	instance application.ComponentInstance
}

type comInstanceLoader struct {
	context    application.Context
	pool       lang.ReleasePool
	cache      map[string]*comInstanceHolder
	todoInject []application.ComponentInstance
	todoInit   []application.ComponentInstance
}

func (inst *comInstanceLoader) init(ctx application.Context, useNewPool bool) application.InstanceContext {

	pool := ctx.GetReleasePool()
	if useNewPool {
		pool = lang.CreateReleasePool()
	}

	inst.context = ctx
	inst.pool = pool
	inst.cache = make(map[string]*comInstanceHolder)
	inst.todoInject = make([]application.ComponentInstance, 0)
	inst.todoInit = make([]application.ComponentInstance, 0)

	return inst
}

func (inst *comInstanceLoader) Pool() lang.ReleasePool {
	return inst.pool
}

func (inst *comInstanceLoader) Context() application.Context {
	return inst.context
}

func (inst *comInstanceLoader) GetString(selector string) (string, error) {
	const prefix = "${"
	const suffix = "}"
	selector = strings.TrimSpace(selector)
	if strings.HasPrefix(selector, prefix) && strings.HasSuffix(selector, suffix) {
		key := selector[len(prefix) : len(selector)-len(suffix)]
		val, err := inst.context.GetProperties().GetPropertyRequired(key)
		if err != nil {
			return "", errors.New("no property for selector: " + selector)
		}
		return val, nil
	}
	return selector, nil
}

func (inst *comInstanceLoader) GetBool(selector string) (bool, error) {
	str, err := inst.GetString(selector)
	if err != nil {
		return false, err
	}
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	b := (str == "true") || (str == "1") || (str == "yes") || (str == "y")
	return b, nil
}

func (inst *comInstanceLoader) getIntXX(selector string, base int, bits int) (int64, error) {
	str, err := inst.GetString(selector)
	if err != nil {
		return 0, err
	}
	base = 0
	n, err := strconv.ParseInt(str, base, bits)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (inst *comInstanceLoader) GetInt(selector string) (int, error) {
	const base = 0
	const bits = 0
	n, err := inst.getIntXX(selector, base, bits)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func (inst *comInstanceLoader) GetInt16(selector string) (int16, error) {
	const base = 10
	const bits = 16
	n, err := inst.getIntXX(selector, base, bits)
	if err != nil {
		return 0, err
	}
	return int16(n), nil
}

func (inst *comInstanceLoader) GetInt32(selector string) (int32, error) {
	const base = 10
	const bits = 32
	n, err := inst.getIntXX(selector, base, bits)
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}

func (inst *comInstanceLoader) GetInt64(selector string) (int64, error) {
	const base = 10
	const bits = 64
	n, err := inst.getIntXX(selector, base, bits)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (inst *comInstanceLoader) addInstance(i application.ComponentInstance) {
	inst.todoInject = append(inst.todoInject, i)
}

func (inst *comInstanceLoader) addInstances(list []application.ComponentInstance) {
	inst.todoInject = append(inst.todoInject, list...)
}

func (inst *comInstanceLoader) getComInstance(h application.ComponentHolder) application.ComponentInstance {
	info := h.GetInfo()
	id := info.GetID()
	h2 := inst.cache[id]
	if h2 != nil {
		return h2.instance
	}
	h2 = &comInstanceHolder{}
	h2.id = id
	h2.holder = h
	h2.instance = h.GetInstance()
	inst.cache[id] = h2
	inst.addInstance(h2.instance)
	return h2.instance
}

func (inst *comInstanceLoader) GetComponent(selector string) (lang.Object, error) {
	holder, err := inst.context.GetComponents().FindComponent(selector)
	if err != nil {
		return nil, err
	}
	ci := inst.getComInstance(holder)
	return ci.Get(), nil
}

func (inst *comInstanceLoader) GetComponents(selector string) ([]lang.Object, error) {
	objlist := make([]lang.Object, 0)
	holders := inst.context.GetComponents().FindComponents(selector)
	for _, h := range holders {
		ci := inst.getComInstance(h)
		objlist = append(objlist, ci.Get())
	}
	return objlist, nil
}

func (inst *comInstanceLoader) GetComponentsByFilter(selector string, f application.ComponentHolderFilter) ([]lang.Object, error) {
	objlist := make([]lang.Object, 0)
	holders := inst.context.GetComponents().FindComponentsWithFilter(selector, f)
	for _, h := range holders {
		ci := inst.getComInstance(h)
		objlist = append(objlist, ci.Get())
	}
	return objlist, nil
}

func (inst *comInstanceLoader) tryInjectOnce(list []application.ComponentInstance) error {
	if list == nil {
		return nil
	}
	for _, item := range list {
		if item.State() >= application.StateInjected {
			continue
		}
		err := item.Inject(inst)
		if err != nil {
			return err
		}
		inst.todoInit = append(inst.todoInit, item)
	}
	return nil
}

func (inst *comInstanceLoader) pushToPool(i application.ComponentInstance) {
	inst.pool.Push(lang.DisposableForFunc(func() error { return i.Destroy() }))
}

func (inst *comInstanceLoader) tryInitOnce(list []application.ComponentInstance) error {
	if list == nil {
		return nil
	}
	for _, item := range list {
		if item.State() >= application.StateInitialled {
			continue
		}
		err := item.Init()
		if err != nil {
			return err
		}
		inst.pushToPool(item)
	}
	return nil
}

func (inst *comInstanceLoader) hasMore(list []application.ComponentInstance) bool {
	if list == nil {
		return false
	}
	return len(list) > 0
}

func (inst *comInstanceLoader) Load() error {

	// inject
	for ttl := 99; ; ttl-- {
		list := inst.todoInject
		inst.todoInject = make([]application.ComponentInstance, 0)
		if inst.hasMore(list) {
			err := inst.tryInjectOnce(list)
			if err != nil {
				return err
			}
		} else {
			break
		}
		if ttl < 1 {
			return errors.New("TTL == 0")
		}
	}

	// init
	return inst.tryInitOnce(inst.todoInit)
}
