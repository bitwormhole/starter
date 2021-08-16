package tests

import (
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

// TestContext 是测试上下文
type TestContext interface {
	TempDir() fs.Path
	Logger() vlog.Logger
	AddCase(c Case)
	AddCaseFunc(fn OnTestFunc)
	CaseManager() CaseManager
	Runner() TestRunner
}
