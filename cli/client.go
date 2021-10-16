package cli

import (
	"context"

	"github.com/bitwormhole/starter/task"
)

// Client 是用于执行命令的同步 API
type Client interface {
	GetContext() context.Context
	GetFactory() ClientFactory

	Execute(cmd string) error
	ExecuteTask(task *Task) error
	ExecuteWithArguments(cmd string, args []string) error

	// execute multi-line commands
	ExecuteScript(script string) error
}

// AsyncClient 是用于执行命令的异步 API
type AsyncClient interface {
	GetContext() context.Context
	GetFactory() ClientFactory

	Execute(cmd string) task.Promise
	ExecuteTask(task *Task) task.Promise
	ExecuteScript(script string) task.Promise
	ExecuteWithArguments(cmd string, args []string) task.Promise
}

// ClientFactory 是命令客户端的工厂
type ClientFactory interface {
	CreateClient(ctx context.Context) Client
	CreateAsyncClient(ctx context.Context) AsyncClient
}
