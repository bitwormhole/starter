package windows

import "github.com/bitwormhole/starter/platforms"

type windowsPlatform struct {
}

func (inst *windowsPlatform) _Impl() platforms.Platform {
	return inst
}

func (inst *windowsPlatform) OS() platforms.OS {
	return inst
}

func (inst *windowsPlatform) Version() string {
	return ""
}

func (inst *windowsPlatform) Name() string {
	return "windows"
}
