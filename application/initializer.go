package application

import (
	"embed"

	"github.com/bitwormhole/starter/collection"
)

type Initializer interface {
	EmbedResources(fs *embed.FS, path string) Initializer
	MountResources(res collection.Resources, path string) Initializer
	SetAttribute(name string, value interface{}) Initializer
	Use(module Module) Initializer
	Run()
}
