package platforms

import (
	"os"
	"runtime"
)

type linuxPlatformFactory struct {
	arch string
}

func (inst *linuxPlatformFactory) _Impl() platformFactory {
	return inst
}

func (inst *linuxPlatformFactory) Create() Platform {
	p := &platformImpl{}
	p.arch = runtime.GOARCH
	p.osName = Linux
	p.osVersion = os.Getenv("OS_VERSION")
	return p
}
