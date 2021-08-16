package tests

import (
	"testing"
)

// ContextForT 函数为 *testing.T 创建一个 TestContext
func ContextForT(t *testing.T) TestContext {

	tc := &DefaultTestContext{}
	tc.Init()

	fsys := tc.MyTempDir.FileSystem()

	tc.MyTempDir = fsys.GetPath(t.TempDir())
	tc.MyRunner = &DefaultRunner{}

	return tc
}
