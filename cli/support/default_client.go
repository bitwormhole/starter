package support

import (
	"context"

	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/task"
)

// DefaultClientFactory 命令客户端工厂
type DefaultClientFactory struct {
	markup.Component `id:"cli-client-factory"`

	CLI *cli.Context `inject:"#cli-context"`
}

func (inst *DefaultClientFactory) _Impl() cli.ClientFactory {
	return inst
}

// CreateClient 创建同步客户端
func (inst *DefaultClientFactory) CreateClient(ctx context.Context) cli.Client {
	service := inst.CLI.Service
	client := &syncClientImpl{
		context: ctx,
		factory: inst,
		service: service,
	}
	return client
}

// CreateAsyncClient 创建异步客户端
func (inst *DefaultClientFactory) CreateAsyncClient(ctx context.Context) cli.AsyncClient {
	service := inst.CLI.Service
	client1 := &syncClientImpl{
		context: ctx,
		factory: inst,
		service: service,
	}
	client2 := &asyncClientImpl{
		sync: client1,
	}
	return client2
}

////////////////////////////////////////////////////////////////////////////////

// 执行命令的客户端
type syncClientImpl struct {
	context context.Context
	factory cli.ClientFactory
	service cli.Service
}

func (inst *syncClientImpl) _Impl() cli.Client {
	return inst
}

func (inst *syncClientImpl) GetFactory() cli.ClientFactory {
	return inst.factory
}

func (inst *syncClientImpl) GetContext() context.Context {
	return inst.context
}

// Execute  todo...
func (inst *syncClientImpl) ExecuteTask(t *cli.Task) error {
	tc := inst.makeTaskContext(t)
	chain := inst.service.GetFilterChain()
	return chain.Handle(tc)
}

// Execute  todo...
func (inst *syncClientImpl) Execute(cmd string) error {
	tc := inst.makeTaskWithLine(cmd)
	return inst.ExecuteTask(tc)
}

// ExecuteWithArguments  todo...
func (inst *syncClientImpl) ExecuteWithArguments(cmd string, args []string) error {
	t := inst.makeTaskWithArgs(cmd, args)
	return inst.ExecuteTask(t)
}

// ExecuteScript  todo...
func (inst *syncClientImpl) ExecuteScript(script string) error {
	t := inst.makeTaskWithScript(script)
	return inst.ExecuteTask(t)
}

func (inst *syncClientImpl) makeTaskWithArgs(cmd string, args []string) *cli.Task {

	builder := cli.TaskListBuilder{}
	builder.AddLine("", 0, args)

	t := &cli.Task{}
	t.Script = ""
	t.TaskList = builder.Create()
	return t
}

func (inst *syncClientImpl) makeTaskWithScript(script string) *cli.Task {

	builder := cli.TaskListBuilder{}
	builder.ParseScript(script)

	t := &cli.Task{}
	t.Script = script
	t.TaskList = builder.Create()
	return t
}

func (inst *syncClientImpl) makeTaskWithLine(command string) *cli.Task {

	builder := cli.TaskListBuilder{}
	builder.ParseScript(command)

	t := &cli.Task{}
	t.Script = command
	t.TaskList = builder.Create()
	return t
}

func (inst *syncClientImpl) makeTaskContext(t *cli.Task) *cli.TaskContext {
	tc := &cli.TaskContext{}
	tc.TaskList = t.TaskList
	return tc
}

////////////////////////////////////////////////////////////////////////////////

type asyncClientImpl struct {
	sync *syncClientImpl
}

func (inst *asyncClientImpl) _Impl() cli.AsyncClient {
	return inst
}

func (inst *asyncClientImpl) GetContext() context.Context {
	return inst.sync.context
}

func (inst *asyncClientImpl) GetFactory() cli.ClientFactory {
	return inst.sync.factory
}

func (inst *asyncClientImpl) Execute(cmd string) task.Promise {
	t := inst.sync.makeTaskWithLine(cmd)
	return inst.ExecuteTask(t)
}

func (inst *asyncClientImpl) ExecuteWithArguments(cmd string, args []string) task.Promise {
	t := inst.sync.makeTaskWithArgs(cmd, args)
	return inst.ExecuteTask(t)
}

func (inst *asyncClientImpl) ExecuteScript(script string) task.Promise {
	t := inst.sync.makeTaskWithScript(script)
	return inst.ExecuteTask(t)
}

func (inst *asyncClientImpl) ExecuteTask(t *cli.Task) task.Promise {

	return task.NewPromise(func(resolve task.ResolveFn, reject task.RejectFn) {
		err := inst.sync.ExecuteTask(t)
		if err == nil {
			resolve(t)
		} else {
			reject(err)
		}
	})
}
