package srctestgo

import (
	"testing"

	"github.com/bitwormhole/starter/starter"
)

func TestMain(t *testing.T) {
	appinit := starter.InitApp()
	appinit.Run()
}
