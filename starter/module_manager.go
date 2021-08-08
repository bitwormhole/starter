package starter

import (
	"errors"
	"sort"

	"github.com/bitwormhole/starter/application"
)

////////////////////////////////////////////////////////////////////////////////

type moduleManager struct {
	table map[string]*moduleHolder
	// list       []*moduleHolder
	depthLimit int
}

func createModuleManager() *moduleManager {
	inst := &moduleManager{}
	inst.init()
	return inst
}

func (inst *moduleManager) init() {
	// inst.list = make([]*moduleHolder, 0)
	inst.table = make(map[string]*moduleHolder)
	inst.depthLimit = 256
}

func (inst *moduleManager) use(mod application.Module, recursion bool) {
	depth := 1
	if recursion {
		inst.useAsRecursion(mod, depth)
	} else {
		inst.tryLoad(mod, depth)
	}
}

func (inst *moduleManager) useAsRecursion(mod application.Module, depth int) error {
	if mod == nil {
		return nil
	}
	if depth > inst.depthLimit {
		return errors.New("the module dependencies tree is too deep")
	}
	inst.tryLoad(mod, depth)
	deps := mod.GetDependencies()
	if deps == nil {
		return nil
	}
	for _, dep := range deps {
		err := inst.useAsRecursion(dep, depth+1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *moduleManager) tryLoad(mod application.Module, depth int) {
	if mod == nil {
		return
	}
	key := inst.keyFor(mod)
	older := inst.table[key]
	if older == nil {
		holder := &moduleHolder{}
		holder.key = key
		holder.module = mod
		holder.depthSum = depth
		inst.table[key] = holder
	} else {
		older.depthSum += depth
	}
}

func (inst *moduleManager) keyFor(mod application.Module) string {
	name := mod.GetName()
	ver := mod.GetVersion()
	return name + "#" + ver
}

func (inst *moduleManager) listAll() []application.Module {
	sorter := &moduleHolderSorter{}
	sorter.init(inst.table)
	return sorter.sort()
}

////////////////////////////////////////////////////////////////////////////////

type moduleHolderSorter struct {
	list []*moduleHolder
}

func (inst *moduleHolderSorter) init(src map[string]*moduleHolder) {
	dst := make([]*moduleHolder, 0)
	if src == nil {
		inst.list = dst
		return
	}
	for _, item := range src {
		dst = append(dst, item)
	}
	inst.list = dst
}

func (inst *moduleHolderSorter) sort() []application.Module {
	sort.Sort(inst)
	src := inst.list
	dst := make([]application.Module, 0)
	for _, item := range src {
		dst = append(dst, item.module)
	}
	return dst
}

func (inst *moduleHolderSorter) Swap(i1, i2 int) {
	o1 := inst.list[i1]
	o2 := inst.list[i2]
	inst.list[i1] = o2
	inst.list[i2] = o1
}

func (inst *moduleHolderSorter) Less(i1, i2 int) bool {
	o1 := inst.list[i1]
	o2 := inst.list[i2]
	return o1.depthSum > o2.depthSum
}

func (inst *moduleHolderSorter) Len() int {
	return len(inst.list)
}

////////////////////////////////////////////////////////////////////////////////

type moduleHolder struct {
	key      string
	module   application.Module
	depthSum int
}

func (inst *moduleHolder) init() {
}

////////////////////////////////////////////////////////////////////////////////
