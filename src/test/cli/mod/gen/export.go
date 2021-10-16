package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigCLITest(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
