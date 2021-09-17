package cli

import "github.com/bitwormhole/starter/task"

// Client 是用于执行命令的 API
type Client interface {
	Execute(command string) task.Promise
	ExecuteWithArguments(args []string) task.Promise
}
