package tests

// OnTestFunc 是测试函数的签名
type OnTestFunc func(ctx TestContext) error

// Case 把测试函数封装成接口
type Case interface {
	OnTest(ctx TestContext) error
}

// CaseInfo 把测试函数封装成接口
type CaseInfo struct {
	ID        string
	Class     string
	Case      Case
	Error     error
	Done      bool
	TimeBegin int64
	TimeEnd   int64
}

// CaseManager 是tests.Case的管理器
type CaseManager interface {
	AddCase(c Case)
	AddCaseFunc(fn OnTestFunc)
	All() []*CaseInfo
}

// TestRunner 是运行测试的入口
type TestRunner interface {
	// application.Looper
	Run(ctx TestContext) error
}
