package etc

import (
	"github.com/bitwormhole/starter/application"
	etcstarter "github.com/bitwormhole/starter/etc/starter"
)

func Config(cb application.ConfigBuilder) error {
	etcstarter.ExportConfig(cb, nil)
	return autoGenConfig(cb)
}
