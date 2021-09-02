package main

import (
	"testing"

	"github.com/bitwormhole/starter/tests"
)

// //go:embed res
// var res embed.FS

func TestEmbedFsRes(t *testing.T) {
	rt, _ := tests.TestingStarter(t).RunEx()
	rt.Loop()
}
