package task

// Controller 任务控制器(用户端接口)
type Controller interface {
	Pause() error
	Resume() error
	Cancel() error
	GetState() State
	GetStatus() Status
	GetError() error
	GetProgress() map[string]*Progress
}
