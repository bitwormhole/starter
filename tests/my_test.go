package tests

import (
	"errors"
	"testing"
)

func TestMyself(t *testing.T) {

	ctx := ContextForT(t)

	dir := ctx.TempDir()
	ctx.Logger().Info("tempDir1=", dir.Path())

	ctx.AddCaseFunc(func(ctx2 TestContext) error {
		dir2 := ctx2.TempDir()
		ctx2.Logger().Info("tempDir2=", dir2.Path())
		return errors.New("bad for test")
	})

	err := ctx.Runner().Run(ctx)
	if err != nil {
		panic(err)
	}
}
