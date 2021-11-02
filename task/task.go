package task

// State 表示任务执行过程中的某个状态
type State int

// Status 表示任务的最终状态
type Status uint

// 定义任务过程中的状态
const (
	StateStarting   = 1
	StateStarted    = 2
	StateRunning    = 3
	StatePaused     = 4
	StateCancelling = 5
	StateStopping   = 6
	StateStopped    = 7
)

// 定义任务的最终状态
const (
	StatusOK        = 200
	StatusCancelled = 300
	StatusFail      = 500
)
