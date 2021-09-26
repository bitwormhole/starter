package cli

import "sort"

////////////////////////////////////////////////////////////////////////////////

type filterChainNode struct {
	next   FilterChain
	filter Filter
}

func (inst *filterChainNode) _Impl() FilterChain {
	return inst
}

func (inst *filterChainNode) Handle(ctx *TaskContext) error {
	n := inst.next
	f := inst.filter
	return f.Handle(ctx, n)
}

////////////////////////////////////////////////////////////////////////////////

type filterChainEnding struct {
}

func (inst *filterChainEnding) _Impl() FilterChain {
	return inst
}

func (inst *filterChainEnding) Handle(ctx *TaskContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type filterRegistration struct {
	priority int
	filter   Filter
}

////////////////////////////////////////////////////////////////////////////////

type filterChainBuilder struct {
	list    []*filterRegistration
	reverse bool
}

func (inst *filterChainBuilder) Len() int {
	return len(inst.list)
}

func (inst *filterChainBuilder) Less(a, b int) bool {
	n1 := inst.list[a].priority
	n2 := inst.list[b].priority
	if inst.reverse {
		return n1 > n2
	}
	return n1 < n2
}

func (inst *filterChainBuilder) Swap(a, b int) {
	item1 := inst.list[a]
	item2 := inst.list[b]
	inst.list[a] = item2
	inst.list[b] = item1
}

func (inst *filterChainBuilder) add(priority int, filter Filter) {
	if filter == nil {
		return
	}
	reg := &filterRegistration{
		priority: priority,
		filter:   filter,
	}
	inst.list = append(inst.list, reg)
}

func (inst *filterChainBuilder) create() FilterChain {
	sort.Sort(inst)
	list := inst.list
	chain := (&filterChainEnding{})._Impl()
	for _, item := range list {
		node := &filterChainNode{
			next:   chain,
			filter: item.filter,
		}
		chain = node
	}
	return chain
}

////////////////////////////////////////////////////////////////////////////////
