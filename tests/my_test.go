package tests

import (
	"testing"

	"github.com/bitwormhole/starter/src/test"
	"github.com/bitwormhole/starter/vlog"
)

func initRes(ti TestingInitializer) {
	ti.UseResources(test.ExportResources())
}

func TestWithDirData(t *testing.T) {

	const path = "dirs/test1"

	i := Starter(t)
	initRes(i)
	rt := i.PrepareTestingDataFromResource(path).RunTest()
	dir := rt.TestingDataDir()

	vlog.Info("testing.data.dir=", dir.Path())
}

func TestWithZipData(t *testing.T) {

	const path = "dirs/test1.zip"

	i := Starter(t)
	initRes(i)
	rt := i.PrepareTestingDataFromResource(path).RunTest()
	dir := rt.TestingDataDir()

	vlog.Info("testing.data.dir=", dir.Path())
}
