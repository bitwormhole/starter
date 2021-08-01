package starter

import (
	"github.com/bitwormhole/starter/application"
	etc "github.com/bitwormhole/starter/etc/starter"
)

func Module() application.Module {
	return etc.Module()
}
