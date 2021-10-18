package platforms

import (
	"os"
	"runtime"
)

type windowsPlatformFactory struct {
	arch string
}

func (inst *windowsPlatformFactory) _Impl() platformFactory {
	return inst
}

func (inst *windowsPlatformFactory) Create() Platform {

	p := &platformImpl{}
	p.arch = runtime.GOARCH
	p.osName = Windows
	p.osVersion = os.Getenv("OS_VERSION")
	return p
}
