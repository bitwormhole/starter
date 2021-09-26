package cli

import "github.com/bitwormhole/starter/task"

// Client 是用于执行命令的 API
type Client interface {
	Execute(task *Task) error

	ExecuteScript(script string) error

	ExecuteWithArguments(args []string) error

	ExecuteAsync(task *Task) task.Promise

	ExecuteScriptAsync(script string) task.Promise

	ExecuteWithArgumentsAsync(args []string) task.Promise
}
