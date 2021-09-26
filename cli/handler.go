package cli

// Handler 是某条具体命令的处理器
type Handler interface {
	Init(service Service) error
	Handle(ctx *TaskContext) error
}

////////////////////////////////////////////////////////////////////////////////

// CommandHelpInfo 是命令的帮助信息
type CommandHelpInfo struct {
	Name        string
	Title       string
	Description string
	Content     string
}

////////////////////////////////////////////////////////////////////////////////

// CommandHelper 是命令的帮助信息提供者, 通常与Handler 实现于同一个结构
type CommandHelper interface {
	GetHelpInfo() *CommandHelpInfo
}
