package application

import (
	"sort"

	"github.com/bitwormhole/starter/lang"
)

const ExitCodeGeneratorClassName = "exit-code-generator"

type ExitCodeGenerator interface {
	ExitCode() int
}

////////////////////////////////////////////////////////////////////////////////

func tryGetExitCodeGenerator(context Context) ExitCodeGenerator {

	list := &exitCoderList{}
	in := context.Injector()
	selector := "." + ExitCodeGeneratorClassName

	in.Inject(selector).AsList().To(func(o lang.Object) bool {
		item, ok := o.(ExitCodeGenerator)
		if ok {
			list.Add(item)
		}
		return ok
	})

	in.Done()
	return list
}

////////////////////////////////////////////////////////////////////////////////
// struct exitCoderItem

type exitCoderItem struct {
	coder    ExitCodeGenerator
	priority int
}

func (inst *exitCoderItem) loadPriority() {
	pp, ok := inst.coder.(lang.PriorityProvider)
	if ok {
		inst.priority = pp.Priority()
	}
}

func (inst *exitCoderItem) tryExit() (int, bool) {
	coder := inst.coder
	if coder == nil {
		return 0, false
	}
	return coder.ExitCode(), true
}

////////////////////////////////////////////////////////////////////////////////
// struct exitCoderList

type exitCoderList struct {
	items []*exitCoderItem
}

func (inst *exitCoderList) _impl_ExitCodeGenerator() ExitCodeGenerator {
	return inst
}

func (inst *exitCoderList) ExitCode() int {
	list := inst.prepare()
	code := 0
	for index := range list {
		item := list[index]
		if item == nil {
			continue
		}
		c, ok := item.tryExit()
		if ok {
			code = c
		}
	}
	return code
}

func (inst *exitCoderList) Add(coder ExitCodeGenerator) {
	if coder == nil {
		return
	}
	list := inst.items
	if list == nil {
		list = []*exitCoderItem{}
	}
	item := &exitCoderItem{coder: coder}
	list = append(list, item)
	inst.items = list
}

func (inst *exitCoderList) prepare() []*exitCoderItem {

	list := inst.items
	if list == nil {
		list = make([]*exitCoderItem, 0)
	}

	list = append(list, &exitCoderItem{}) // for debug

	for index := range list {
		list[index].loadPriority()
	}

	inst.items = list
	sort.Sort(inst)
	return list
}

func (inst *exitCoderList) Len() int {
	return len(inst.items)
}

func (inst *exitCoderList) Less(ia int, ib int) bool {
	a := inst.items[ia].priority
	b := inst.items[ib].priority
	return a < b
}

func (inst *exitCoderList) Swap(ia int, ib int) {
	a := inst.items[ia]
	b := inst.items[ib]
	inst.items[ia] = b
	inst.items[ib] = a
}

////////////////////////////////////////////////////////////////////////////////
// EOF
