package platforms

import (
	"os"
	"runtime"
)

type osxPlatformFactory struct {
	arch string
}

func (inst *osxPlatformFactory) _Impl() platformFactory {
	return inst
}

func (inst *osxPlatformFactory) Create() Platform {
	p := &platformImpl{}
	p.arch = runtime.GOARCH
	p.osName = MacOS
	p.osVersion = os.Getenv("OS_VERSION")
	return p
}
