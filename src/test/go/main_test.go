package main

import (
	"testing"

	"github.com/bitwormhole/starter"
)

func TestMain(t *testing.T) {
	appinit := starter.InitApp()
	appinit.Run()
}
