package tests

import (
	"os"
	"strconv"
	"time"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

// DefaultTestContext 是默认的测试上下文
type DefaultTestContext struct {
	MyCaseManager CaseManager
	MyLogger      vlog.Logger
	MyRunner      TestRunner
	MyTempDir     fs.Path
}

func (inst *DefaultTestContext) _Impl() TestContext {
	return inst
}

// Init 执行默认的初始化过程
func (inst *DefaultTestContext) Init() {

	now := time.Now().Unix() * 1000

	inst.MyLogger = vlog.Default()
	inst.MyTempDir = fs.Default().GetPath(os.TempDir()).GetChild("starter/tests/t" + strconv.FormatInt(now, 10))
	inst.MyCaseManager = &DefaultCaseManager{}
	inst.MyRunner = &DefaultRunner{}
}

// Init 执行默认的初始化过程
func (inst *DefaultTestContext) InitWith(ctx TestContext) {
	inst.MyLogger = ctx.Logger()
	inst.MyTempDir = ctx.TempDir()
	inst.MyCaseManager = ctx.CaseManager()
	inst.MyRunner = ctx.Runner()
}

// Clone 创建上下文的副本
func (inst *DefaultTestContext) Clone() TestContext {
	tc2 := &DefaultTestContext{}
	tc2.MyCaseManager = inst.MyCaseManager
	tc2.MyLogger = inst.MyLogger
	tc2.MyRunner = inst.MyRunner
	tc2.MyTempDir = inst.MyTempDir
	return tc2
}

// TempDir 当前测试的临时文件夹
func (inst *DefaultTestContext) TempDir() fs.Path {
	return inst.MyTempDir
}

// Logger 日志接口
func (inst *DefaultTestContext) Logger() vlog.Logger {
	return inst.MyLogger
}

// CaseManager 测试用例管理器
func (inst *DefaultTestContext) CaseManager() CaseManager {
	return inst.MyCaseManager
}

// Runner 测试执行器
func (inst *DefaultTestContext) Runner() TestRunner {
	return inst.MyRunner
}

// AddCase 添加一个用例到上下文
func (inst *DefaultTestContext) AddCase(c Case) {
	inst.MyCaseManager.AddCase(c)
}

// AddCaseFunc 添加一个用例到上下文
func (inst *DefaultTestContext) AddCaseFunc(fn OnTestFunc) {
	inst.MyCaseManager.AddCaseFunc(fn)
}
