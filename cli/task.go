package cli

import (
	"context"
)

// TaskUnit 表示一条简单的命令
type TaskUnit struct {
	LineNumber  int
	CommandName string
	CommandLine string
	Arguments   []string
}

// Task 表示将要执行的任务
type Task struct {
	Context  context.Context
	TaskList []*TaskUnit
	Script   string
}

// TaskContext 表示正在执行的任务
type TaskContext struct {
	Context     context.Context
	CurrentTask *TaskUnit
	Handler     Handler
	Service     Service
	TaskList    []*TaskUnit
	Console     Console
}

// Clone 生成 TaskContext 的副本（浅拷贝）
func (inst *TaskContext) Clone() *TaskContext {
	child := &TaskContext{
		Context:     inst.Context,
		CurrentTask: inst.CurrentTask,
		Handler:     inst.Handler,
		Service:     inst.Service,
		TaskList:    inst.TaskList,
		Console:     inst.Console,
	}
	return child
}
