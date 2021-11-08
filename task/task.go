package task

// State 表示任务执行过程中的某个状态
type State int

// Status 表示任务的最终状态
type Status uint

// 定义任务过程中的状态
const (
	StateStarting   State = 1
	StateStarted    State = 2
	StateRunning    State = 3
	StatePaused     State = 4
	StateCancelling State = 5
	StateStopping   State = 6
	StateStopped    State = 7
)

// 定义任务的最终状态
const (
	StatusOK        Status = 200
	StatusCancelled Status = 300
	StatusFail      Status = 500
)
