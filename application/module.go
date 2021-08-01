package application

type Module interface {
	GetName() string
	GetVersion() string
	GetRevision() int
	GetDependencies() []Module
	Apply(cb ConfigBuilder) error
}

type ModuleApplyFunc func(cb ConfigBuilder) error

type ModuleDefine struct {
	Name         string
	Version      string
	Revision     int
	Dependencies []Module
	HandleApply  ModuleApplyFunc
}

func (inst *ModuleDefine) __impl__() Module {
	return inst
}

func (inst *ModuleDefine) GetDependencies() []Module {
	src := inst.Dependencies
	dst := make([]Module, 0)
	if src == nil {
		return dst
	}
	for index := range src {
		mod := src[index]
		if mod != nil {
			dst = append(dst, mod)
		}
	}
	return dst
}

func (inst *ModuleDefine) GetName() string {
	return inst.Name
}

func (inst *ModuleDefine) GetRevision() int {
	return inst.Revision
}

func (inst *ModuleDefine) GetVersion() string {
	return inst.Version
}

func (inst *ModuleDefine) Apply(cb ConfigBuilder) error {
	return inst.HandleApply(cb)
}
