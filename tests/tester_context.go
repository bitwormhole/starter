package tests

import "github.com/bitwormhole/starter/application"

// ContextForApp 函数为 application.Context 创建一个 TestContext
func ContextForApp(ac application.Context) TestContext {
	tc := &DefaultTestContext{}
	tc.Init()
	tc.MyRunner = &DefaultRunner{}
	return tc
}
