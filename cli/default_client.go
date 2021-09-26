package cli

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/task"
)

// DefaultClient 执行命令的客户端
type DefaultClient struct {
	markup.Component `id:"cli-client" class:"cli-client"`
	Service          Service `inject:"#cli-service"`
}

func (inst *DefaultClient) _Impl() Client {
	return inst
}

// Execute  todo...
func (inst *DefaultClient) Execute(t *Task) error {
	tc := inst.makeTaskContext(t)
	chain := inst.Service.GetFilterChain()
	return chain.Handle(tc)
}

// ExecuteScript  todo...
func (inst *DefaultClient) ExecuteScript(command string) error {
	t := inst.makeTaskWithScript(command)
	return inst.Execute(t)
}

// ExecuteWithArguments  todo...
func (inst *DefaultClient) ExecuteWithArguments(args []string) error {
	t := inst.makeTaskWithArgs(args)
	return inst.Execute(t)
}

// ExecuteAsync  todo...
func (inst *DefaultClient) ExecuteAsync(t *Task) task.Promise {
	return task.NewPromise(func(resolve task.ResolveFn, reject task.RejectFn) {
		err := inst.Execute(t)
		if err == nil {
			resolve(t)
		} else {
			reject(err)
		}
	})
}

// ExecuteScriptAsync  todo...
func (inst *DefaultClient) ExecuteScriptAsync(command string) task.Promise {
	t := inst.makeTaskWithScript(command)
	return inst.ExecuteAsync(t)
}

// ExecuteWithArgumentsAsync  todo...
func (inst *DefaultClient) ExecuteWithArgumentsAsync(args []string) task.Promise {
	t := inst.makeTaskWithArgs(args)
	return inst.ExecuteAsync(t)
}

func (inst *DefaultClient) makeTaskWithArgs(args []string) *Task {

	builder := &taskListBuilder{}
	builder.addLine("", 0, args)

	t := &Task{}
	t.Script = ""
	t.TaskList = builder.create()
	return t
}

func (inst *DefaultClient) makeTaskWithScript(script string) *Task {

	builder := &taskListBuilder{}
	builder.parseScript(script)

	t := &Task{}
	t.Script = script
	t.TaskList = builder.create()
	return t
}

func (inst *DefaultClient) makeTaskContext(t *Task) *TaskContext {
	tc := &TaskContext{}
	tc.TaskList = t.TaskList
	return tc
}
