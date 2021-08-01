package srctestgo

import (
	"testing"

	starter "github.com/bitwormhole/starter"
)

func TestMain(t *testing.T) {
	appinit := starter.Init().MountResources(starter.GetResources(), "/")
	appinit.Use(starter.Module())
	appinit.Run()
}
