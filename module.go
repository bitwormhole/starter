package starter

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/etc/starter"
)

func Module() application.Module {
	return &application.DefineModule{
		Name:     "github.com/bitwormhole/starter",
		Version:  "1.0",
		Revision: 1,
		OnMount:  func(cb application.ConfigBuilder) error { return starter.Config(cb) },
	}
}
