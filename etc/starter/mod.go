package starter

import "github.com/bitwormhole/starter/application"

func Module() application.Module {
	return &application.ModuleDefine{
		Name:     "github.com/bitwormhole/starter/etc/starter",
		Version:  "1.0",
		Revision: 1,
		HandleApply: func(cb application.ConfigBuilder) error {
			return Config(cb)
		},
	}
}
