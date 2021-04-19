package runtime

import (
	"errors"
	"sort"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct

type componentLoading struct {
	holder       application.ComponentHolder
	instance     application.ComponentInstance
	loadingOrder int
}

type creationComponentLoader struct {
	core              *creationContextCore
	loadingOrderCount int
}

type runtimeComponentLoader struct {
	core *runtimeContextCore
}

type componentLoadingSorter struct {
	items []*componentLoading
}

////////////////////////////////////////////////////////////////////////////////
// impl creationComponentLoader

func (inst *creationComponentLoader) createNewLoading(holder application.ComponentHolder) *componentLoading {
	instance := holder.GetInstance()
	loading := &componentLoading{}
	loading.instance = instance
	loading.holder = holder
	loading.loadingOrder = inst.loadingOrderCount
	inst.loadingOrderCount++
	return loading
}

func (inst *creationComponentLoader) loadComponent(holder application.ComponentHolder) (lang.Object, error) {

	if holder == nil {
		return nil, errors.New("holder==nil:ComponentHolder")
	}

	id := holder.GetInfo().GetID()
	cache := inst.core.cache
	loading := cache[id]

	if loading == nil {
		loading = inst.createNewLoading(holder)
		cache[id] = loading
		// do inject
		ctx := inst.core.proxy
		loading.instance.Inject(ctx)
	}

	// result
	target := loading.instance.Get()
	return target, nil
}

func (inst *creationComponentLoader) loadComponents(holders []application.ComponentHolder) ([]lang.Object, error) {
	if holders == nil {
		return nil, errors.New("holders==nil:ComponentHolder")
	}
	dst := make([]lang.Object, 0)
	for index := range holders {
		h := holders[index]
		target, err := inst.loadComponent(h)
		if err != nil {
			return nil, err
		}
		if target == nil {
			continue
		}
		dst = append(dst, target)
	}
	return dst, nil
}

func (inst *creationComponentLoader) startAllComponents() error {

	table := inst.core.cache
	list := make([]*componentLoading, 0)
	for key := range table {
		list = append(list, table[key])
	}

	// todo: sort by loadingOrder
	sort.Sort(&componentLoadingSorter{items: list})

	for index := range list {
		item := list[index]
		err := item.instance.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// impl runtimeComponentLoader

func (inst *runtimeComponentLoader) loadComponent(holder application.ComponentHolder) (lang.Object, error) {

	if holder == nil {
		return nil, errors.New("holder==nil:ComponentHolder")
	}

	info := holder.GetInfo()
	scope := info.GetScope()

	if scope == application.ScopeSingleton {
		instance := holder.GetInstance()
		if instance.IsLoaded() {
			target := instance.Get()
			return target, nil
		}
	}

	// do loading
	ctx := inst.core.context
	cc := ctx.OpenCreationContext(scope)
	id := info.GetID()
	components := cc.GetContext().GetComponents()

	target, err := components.GetComponent(id)
	if err != nil {
		return nil, err
	}

	err = cc.Close()
	if err != nil {
		return nil, err
	}

	return target, nil
}

func (inst *runtimeComponentLoader) loadComponents(holders []application.ComponentHolder) ([]lang.Object, error) {
	if holders == nil {
		return nil, errors.New("holders==nil:ComponentHolder")
	}
	dst := make([]lang.Object, 0)
	for index := range holders {
		h := holders[index]
		target, err := inst.loadComponent(h)
		if err != nil {
			return nil, err
		}
		if target == nil {
			continue
		}
		dst = append(dst, target)
	}
	return dst, nil
}

////////////////////////////////////////////////////////////////////////////////
// impl componentLoadingSorter

func (inst *componentLoadingSorter) Len() int {
	return len(inst.items)
}

func (inst *componentLoadingSorter) Less(i, j int) bool {
	a := inst.items[i]
	b := inst.items[j]
	return a.loadingOrder > b.loadingOrder
}

func (inst *componentLoadingSorter) Swap(i, j int) {
	a := inst.items[i]
	b := inst.items[j]
	inst.items[i] = b
	inst.items[j] = a
}

////////////////////////////////////////////////////////////////////////////////
// EOF
