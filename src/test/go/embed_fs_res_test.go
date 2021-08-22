package main

import (
	"testing"

	"github.com/bitwormhole/starter"
)

// //go:embed res
// var res embed.FS

func TestEmbedFsRes(t *testing.T) {
	starter.InitApp().Run()
}
