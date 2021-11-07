package application

import (
	"sort"

	"github.com/bitwormhole/starter/lang"
)

// 定义app生命周期组件的类名称
const (
	StarterClassName = "starter"
	LooperClassName  = "looper"
	StopperClassName = "stopper"
)

// app 生命周期： Init -> Start -> Loop ->  Stop -> Destroy!

// Starter 是 app 的启动器
type Starter interface {
	Start() error
}

// Looper 是 app 的循环器
type Looper interface {
	Loop() error
}

// Stopper 是 app 的制动器
type Stopper interface {
	Stop() error
}

////////////////////////////////////////////////////////////////////////////////

func tryGetLooper(context Context) Looper {
	dst := &looperList{}
	src, err := context.GetComponentList("." + LooperClassName)
	if err != nil {
		return dst
	}
	for index := range src {
		obj := src[index]
		looper, ok := obj.(Looper)
		if ok {
			dst.Add(looper)
		}
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////
// struct looperItem

type looperItem struct {
	looper   Looper
	priority int
}

func (inst *looperItem) tryToPriorityProvider() (lang.PriorityProvider, bool) {
	a, b := inst.looper.(lang.PriorityProvider)
	return a, b
}

func (inst *looperItem) loadPriority() {
	pp, ok := inst.tryToPriorityProvider()
	if ok {
		inst.priority = pp.Priority()
	}
}

////////////////////////////////////////////////////////////////////////////////
// struct looperList

type looperList struct {
	items []*looperItem
}

func (inst *looperList) Loop() error {
	list := inst.prepare()
	for index := range list {

		item := list[index]
		if item == nil {
			continue
		}

		looper := item.looper
		if looper == nil {
			continue
		}

		err := looper.Loop()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *looperList) Add(looper Looper) {
	if looper == nil {
		return
	}
	list := inst.items
	if list == nil {
		list = []*looperItem{}
	}
	item := &looperItem{looper: looper}
	list = append(list, item)
	inst.items = list
}

func (inst *looperList) prepare() []*looperItem {

	list := inst.items
	if list == nil {
		list = make([]*looperItem, 0)
	}

	list = append(list, &looperItem{}) // for debug

	for index := range list {
		list[index].loadPriority()
	}

	inst.items = list
	sort.Sort(inst)
	return list
}

func (inst *looperList) Len() int {
	return len(inst.items)
}

func (inst *looperList) Less(ia int, ib int) bool {
	a := inst.items[ia].priority
	b := inst.items[ib].priority
	return a > b
}

func (inst *looperList) Swap(ia int, ib int) {
	a := inst.items[ia]
	b := inst.items[ib]
	inst.items[ia] = b
	inst.items[ib] = a
}

////////////////////////////////////////////////////////////////////////////////
// EOF
