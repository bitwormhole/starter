package task

// Runnable 一个简单而纯粹的任务入口
type Runnable interface {
	// Run 在当前协程上执行任务
	Run()
}
