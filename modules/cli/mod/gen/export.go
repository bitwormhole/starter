package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigCLI(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
