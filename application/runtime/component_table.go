package runtime

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter/application"
)

type componentTable struct {
	table   map[string]application.ComponentHolder
	context application.Context
}

func (inst *componentTable) init(ctx application.Context) application.Components {
	inst.context = ctx
	inst.table = make(map[string]application.ComponentHolder)
	return inst
}

func (inst *componentTable) GetComponentNameList(includeAliases bool) []string {
	dst := make([]string, 0)
	src := inst.table
	if src == nil {
		return dst
	}
	for key := range src {
		item := src[key]
		if item == nil {
			continue
		}
		if !includeAliases {
			comId := item.GetInfo().GetID()
			if key != comId {
				continue
			}
		}
		dst = append(dst, key)
	}
	return dst
}

func (inst *componentTable) GetComponent(selector string) (application.ComponentHolder, error) {

	selector = strings.TrimSpace(selector)

	if strings.HasPrefix(selector, "#") {
		// id
		id := selector[1:]
		return inst.findComponentByID(id)
	} else if strings.HasPrefix(selector, ".") {
		// class
		class := selector[1:]
		items := inst.findComponentsByClass(class)
		return inst.getOnlyOneFromList(items, selector)
	}
	return nil, errors.New("cannot find component with selector: " + selector)
}

func (inst *componentTable) GetComponents(selector string) []application.ComponentHolder {

	selector = strings.TrimSpace(selector)

	if strings.HasPrefix(selector, "#") {
		// id
		id := selector[1:]
		item, err := inst.findComponentByID(id)
		if err == nil {
			return []application.ComponentHolder{item}
		} else {
			return []application.ComponentHolder{}
		}
	} else if strings.HasPrefix(selector, ".") {
		// class
		class := selector[1:]
		return inst.findComponentsByClass(class)
	} else if selector == "*.scope(singleton)" {
		// scope: singleton
		return inst.GetComponentsByFilter(func(name string, ch application.ComponentHolder) bool {
			if ch == nil {
				return false
			}
			scope := ch.GetInfo().GetScope()
			if scope != application.ScopeSingleton {
				return false
			}
			comId := ch.GetInfo().GetID()
			return (comId == name)
		})
	} else if selector == "*" {
		// all (without aliases)
		return inst.GetComponentsByFilter(func(name string, ch application.ComponentHolder) bool {
			if ch == nil {
				return false
			}
			comId := ch.GetInfo().GetID()
			return (comId == name)
		})
	}
	return []application.ComponentHolder{}
}

func (inst *componentTable) findComponentByID(id string) (application.ComponentHolder, error) {
	item := inst.table[id]
	if item == nil {
		return nil, errors.New("no component with name: " + id)
	}
	return item, nil
}

func (inst *componentTable) getOnlyOneFromList(list []application.ComponentHolder, selector string) (application.ComponentHolder, error) {
	if list == nil {
		return nil, errors.New("results==nil")
	}
	size := len(list)
	if size == 1 {
		item := list[0]
		if item != nil {
			return item, nil
		}
	} else if size > 1 {
		return nil, errors.New("more than one components are marked with class: " + selector)
	}
	return nil, errors.New("no component with class: " + selector)
}

func (inst *componentTable) findComponentsByClass(className string) []application.ComponentHolder {
	// by filter
	return inst.GetComponentsByFilter(func(name string, ch application.ComponentHolder) bool {
		if ch == nil {
			return false
		}
		all := ch.GetInfo().GetClasses()
		if all == nil {
			return false
		}
		for index := range all {
			if all[index] == className {
				return true
			}
		}
		return false
	})
}

func (inst *componentTable) GetComponentsByFilter(filter application.ComponentHolderFilter) []application.ComponentHolder {
	src := inst.table
	dst := make([]application.ComponentHolder, 0)
	if src == nil {
		return dst
	}
	if filter == nil {
		filter = func(name string, holder application.ComponentHolder) bool {
			return true
		}
	}
	for key := range src {
		item := src[key]
		if item == nil {
			continue
		}
		if item.GetInfo().GetID() != key {
			continue
		}
		if !filter(key, item) {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

func (inst *componentTable) Export(dst map[string]application.ComponentHolder) map[string]application.ComponentHolder {
	if dst == nil {
		dst = make(map[string]application.ComponentHolder)
	}
	src := inst.table
	if src == nil {
		return dst
	}
	for key := range src {
		dst[key] = src[key]
	}
	return dst
}

func (inst *componentTable) Import(src map[string]application.ComponentHolder) {
	if src == nil {
		return
	}
	context := inst.context
	dst := inst.table
	if dst == nil {
		dst = make(map[string]application.ComponentHolder)
	}
	for key := range src {
		old := dst[key]
		if old != nil {
			// skip
			continue
		}
		item := src[key]
		if item == nil {
			continue
		}
		if item.GetInfo().GetID() != key {
			// key is alias, skip
			continue
		}
		item = item.MakeChild(context)
		// for aliases
		aliases := item.GetInfo().GetAliases()
		for index := range aliases {
			alias := aliases[index]
			dst[alias] = item
		}
		dst[key] = item
	}
	inst.table = dst
}
