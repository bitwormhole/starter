package task

// Progress 进度对象
type Progress struct {
	TaskID    string // 任务ID
	Name      string // 项目名称(ID)
	Title     string // 标题
	Unit      string // 单位
	Value     int64  // 当前值
	ValueMin  int64  // 最小值
	ValueMax  int64  // 最大值
	Done      bool   // 是否已完成
	Cancelled bool   // 是否已取消
}

// ProgressControlHandlerFn 进度控制处理函数
type ProgressControlHandlerFn func(reporter ProgressReporter) error

// ProgressReporter 进度报告者（服务端接口）
type ProgressReporter interface {
	Report(p *Progress)
	UpdateStatus(s Status)
	UpdateState(s State)

	HandleCancel(f ProgressControlHandlerFn)
	HandlePause(f ProgressControlHandlerFn)
	HandleResume(f ProgressControlHandlerFn)
}
