package platforms

import "runtime"

type platformFactory interface {
	Create() Platform
}

////////////////////////////////////////////////////////////////////////////////

type platformImpl struct {
	arch      string
	osName    string
	osVersion string
}

func (inst *platformImpl) _Impl() Platform {
	return inst
}

func (inst *platformImpl) GetOS() OS {
	return inst
}

func (inst *platformImpl) OS() string {
	return inst.osName
}

func (inst *platformImpl) Arch() string {
	return inst.arch
}

func (inst *platformImpl) Version() string {
	return inst.osVersion
}

func (inst *platformImpl) Name() string {
	return inst.osName
}

////////////////////////////////////////////////////////////////////////////////

func initCurrent() Platform {

	// arch := runtime.GOARCH
	osName := runtime.GOOS
	var factory platformFactory = nil

	switch osName {

	case "windows":
		factory = &pFactoryWindows{}
		break

	case "linux":
		factory = &pFactoryLinux{}
		break

	case "darwin":
		factory = &pFactoryOSX{}
		break

	case "freebsd":
		break
	default:
		break
	}

	if factory == nil {
		panic("unsupported os: " + osName)
	}

	return factory.Create()
}
