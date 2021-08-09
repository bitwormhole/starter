package osx

import "github.com/bitwormhole/starter/platforms"

type osxPlatform struct {
}

func (inst *osxPlatform) _Impl() platforms.Platform {
	return inst
}

func (inst *osxPlatform) OS() platforms.OS {
	return inst
}

func (inst *osxPlatform) Version() string {
	return ""
}

func (inst *osxPlatform) Name() string {
	return "osx"
}
