package platforms

import (
	"os"
	"runtime"
)

type pFactoryWindows struct {
	arch string
}

func (inst *pFactoryWindows) _Impl() platformFactory {
	return inst
}

func (inst *pFactoryWindows) Create() Platform {

	p := &platformImpl{}
	p.arch = runtime.GOARCH
	p.osName = Windows
	p.osVersion = os.Getenv("OS_VERSION")
	return p
}
