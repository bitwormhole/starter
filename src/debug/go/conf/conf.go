package conf

import (
	"github.com/bitwormhole/starter/application"
)

func ExportConfig(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
