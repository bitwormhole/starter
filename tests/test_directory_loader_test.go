package tests

import (
	"testing"

	srctest "github.com/bitwormhole/starter/src/test"
	"github.com/bitwormhole/starter/vlog"
)

func TestTheLoaderLoadFolder(t *testing.T) {

	ctx := ContextForT(t)

	res := srctest.ExportResources()
	loader := &TestDirectoryLoader{}
	loader.Init(res, ctx)

	dir, err := loader.LoadFromFolder("dirs/test1")
	if err != nil {
		t.Error(err)
		return
	}

	vlog.Debug("base=", dir.Path())
}

func TestTheLoaderLoadZip(t *testing.T) {

	ctx := ContextForT(t)

	res := srctest.ExportResources()
	loader := &TestDirectoryLoader{}
	loader.Init(res, ctx)

	dir, err := loader.LoadFromZipFile("dirs/test1.zip")
	if err != nil {
		t.Error(err)
	}

	vlog.Debug("base=", dir.Path())
}
