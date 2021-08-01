package application

type Module interface {
	GetName() string
	GetVersion() string
	GetRevision() int
	GetDependencies() []Module
	Apply(cb ConfigBuilder) error
}

type OnMountFunc func(cb ConfigBuilder) error

type DefineModule struct {
	Name         string
	Version      string
	Revision     int
	Dependencies []Module
	OnMount      OnMountFunc
}

func (inst *DefineModule) __impl__() Module {
	return inst
}

func (inst *DefineModule) GetDependencies() []Module {
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

func (inst *DefineModule) GetName() string {
	return inst.Name
}

func (inst *DefineModule) GetRevision() int {
	return inst.Revision
}

func (inst *DefineModule) GetVersion() string {
	return inst.Version
}

func (inst *DefineModule) Apply(cb ConfigBuilder) error {
	return inst.OnMount(cb)
}
