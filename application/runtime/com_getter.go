package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////

type ComGetter struct {
	context application.Context
}

func (inst *ComGetter) Init(ctx application.Context) {
	inst.context = ctx
}

func (inst *ComGetter) GetOne(selector string) (lang.Object, error) {
	// find
	ctx := inst.context
	holder, err := ctx.GetComponents().GetComponent(selector)
	if err != nil {
		return nil, err
	}
	if inst.isComponentReady(holder) {
		com := holder.GetInstance().Get()
		return com, nil
	}
	// load
	loading, err := ctx.ComponentLoader().OpenLoading(ctx)
	if err != nil {
		return nil, err
	}
	defer loading.Close()
	com, err := loading.Load(holder)
	if err != nil {
		return nil, err
	}
	err = loading.Close()
	if err != nil {
		return nil, err
	}
	return com, nil
}

func (inst *ComGetter) GetList(selector string) ([]lang.Object, error) {
	// find
	ctx := inst.context
	dstlist := make([]lang.Object, 0)
	holders := inst.context.GetComponents().GetComponents(selector)
	if inst.isAllComponentsReady(holders) {
		for index := range holders {
			holder := holders[index]
			com := holder.GetInstance().Get()
			dstlist = append(dstlist, com)
		}
		return dstlist, nil
	}
	// load
	loading, err := ctx.ComponentLoader().OpenLoading(ctx)
	if err != nil {
		return nil, err
	}
	defer loading.Close()
	comlist, err := loading.LoadAll(holders)
	if err != nil {
		return nil, err
	}
	err = loading.Close()
	if err != nil {
		return nil, err
	}
	return comlist, nil
}

func (inst *ComGetter) isComponentReady(h application.ComponentHolder) bool {
	if h == nil {
		return false
	}
	info := h.GetInfo()
	scope := info.GetScope()
	if scope == application.ScopeSingleton {
		instance := h.GetInstance()
		return instance.IsLoaded()
	} else if scope == application.ScopePrototype {
		return false
	} else {
		return false
	}
}

func (inst *ComGetter) isAllComponentsReady(list []application.ComponentHolder) bool {
	if list == nil {
		return true
	}
	for index := range list {
		item := list[index]
		if item == nil {
			continue
		}
		if !inst.isComponentReady(item) {
			return false
		} else {
			continue
		}
	}
	return true
}
