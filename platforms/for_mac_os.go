package platforms

import (
	"os"
	"runtime"
)

type pFactoryOSX struct {
	arch string
}

func (inst *pFactoryOSX) _Impl() platformFactory {
	return inst
}

func (inst *pFactoryOSX) Create() Platform {
	p := &platformImpl{}
	p.arch = runtime.GOARCH
	p.osName = MacOS
	p.osVersion = os.Getenv("OS_VERSION")
	return p
}
