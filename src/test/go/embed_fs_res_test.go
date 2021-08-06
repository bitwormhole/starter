package srctestgo

import (
	"embed"
	"testing"

	"github.com/bitwormhole/starter/starter"
)

//go:embed res
var res embed.FS

func TestEmbedFsRes(t *testing.T) {
	starter.InitApp().EmbedResources(&res, "res").Run()
}
