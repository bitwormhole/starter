package platforms

import (
	"os"
	"runtime"
)

type pFactoryLinux struct {
	arch string
}

func (inst *pFactoryLinux) _Impl() platformFactory {
	return inst
}

func (inst *pFactoryLinux) Create() Platform {
	p := &platformImpl{}
	p.arch = runtime.GOARCH
	p.osName = Linux
	p.osVersion = os.Getenv("OS_VERSION")
	return p
}
