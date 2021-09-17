package task

// Executor 接口表示一个任务的执行者
type Executor interface {
	Execute(r Runnable)
}
