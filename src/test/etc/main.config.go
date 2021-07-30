package etc

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/etc/starter"
)

func Config(cb application.ConfigBuilder) error {

	starter.Config(cb)

	return autoGenConfig(cb)
}
