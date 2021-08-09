package linux

import "github.com/bitwormhole/starter/platforms"

type linuxPlatform struct {
}

func (inst *linuxPlatform) _Impl() platforms.Platform {
	return inst
}

func (inst *linuxPlatform) OS() platforms.OS {
	return inst
}

func (inst *linuxPlatform) Version() string {
	return ""
}

func (inst *linuxPlatform) Name() string {
	return "osx"
}
