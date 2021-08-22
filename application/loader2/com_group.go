package loader2

import "github.com/bitwormhole/starter/application"

type componentGroup struct {
	name    string                        //  aka selector like '#id', '.class', '*'
	holders []application.ComponentHolder // 分组成员
}

func (inst *componentGroup) _Impl() application.ComponentGroup {
	return inst
}

func (inst *componentGroup) init(name string) application.ComponentGroup {
	inst.name = name
	inst.holders = make([]application.ComponentHolder, 0)
	return inst
}

func (inst *componentGroup) add(h application.ComponentHolder) application.ComponentGroup {
	inst.holders = append(inst.holders, h)
	return inst
}

func (inst *componentGroup) Size() int {
	list := inst.holders
	if list == nil {
		return 0
	}
	return len(list)
}

func (inst *componentGroup) ListAll() []application.ComponentHolder {
	src := inst.holders
	if src == nil {
		return []application.ComponentHolder{}
	}
	dst := make([]application.ComponentHolder, 0, len(src))
	for _, item := range src {
		dst = append(dst, item)
	}
	return dst
}

func (inst *componentGroup) ListWithFilter(f application.ComponentHolderFilter) []application.ComponentHolder {
	src := inst.holders
	if src == nil {
		return []application.ComponentHolder{}
	}
	dst := make([]application.ComponentHolder, 0, 4)
	for _, item := range src {
		id := item.GetInfo().GetID()
		if f(id, item) {
			dst = append(dst, item)
		}
	}
	return dst
}

func (inst *componentGroup) Name() string {
	return inst.name
}
