package srctestgo

import (
	"testing"

	"github.com/bitwormhole/starter"
)

func TestMain(t *testing.T) {
	ai := starter.Init().MountResources(starter.GetResources(), "/")
	ai.Use(starter.Module())
	ai.Run()
}
