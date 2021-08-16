package tests

type DefaultCaseManager struct {
	items []*CaseInfo
}

func (inst *DefaultCaseManager) _Impl() CaseManager {
	return inst
}

func (inst *DefaultCaseManager) addItem(item *CaseInfo) {
	if item == nil {
		return
	}
	list := inst.items
	if list == nil {
		list = make([]*CaseInfo, 0)
	}
	list = append(list, item)
	inst.items = list
}

func (inst *DefaultCaseManager) AddCase(c Case) {
	info := &CaseInfo{}
	info.Case = c
	info.ID = ""
	info.Class = ""
	inst.addItem(info)
}

func (inst *DefaultCaseManager) AddCaseFunc(fn OnTestFunc) {
	info := &CaseInfo{}
	info.Case = &caseForFunc{fn: fn}
	info.ID = ""
	info.Class = ""
	inst.addItem(info)
}

func (inst *DefaultCaseManager) All() []*CaseInfo {
	src := inst.items
	dst := make([]*CaseInfo, 0)
	if src == nil {
		return dst
	}
	for _, item := range src {
		dst = append(dst, item)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type caseForFunc struct {
	fn OnTestFunc
}

func (inst *caseForFunc) _Impl() Case {
	return inst
}

func (inst *caseForFunc) OnTest(ctx TestContext) error {
	return inst.fn(ctx)
}
