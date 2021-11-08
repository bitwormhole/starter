package task

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/vlog"
)

// Progress 进度对象
type Progress struct {
	TaskID  string // 任务ID
	Name    string // 项目名称(ID)
	Title   string // 标题
	Unit    string // 单位
	Message string // 描述当前进度的消息

	Value     int64 // 当前值
	ValueMin  int64 // 最小值
	ValueMax  int64 // 最大值
	TimeBegin int64 // 开始时间
	TimeEnd   int64 // 结束时间

	// Done      bool // 是否已完成 (已废弃，用 State)
	// Cancelled bool // 是否已取消（已废弃，用 Statue）
	State  State  // 过程中的状态
	Status Status // 最终状态
}

// ProgressControlHandlerFn 进度控制处理函数
type ProgressControlHandlerFn func(reporter ProgressReporter) error

// ProgressReporter 进度报告者（服务端接口）
type ProgressReporter interface {
	Report(p *Progress)

	// Update 更新状态 state|status
	Update(p *Progress)

	HandleCancel(f ProgressControlHandlerFn)
	HandlePause(f ProgressControlHandlerFn)
	HandleResume(f ProgressControlHandlerFn)
}

// ProgressReporterFactory 进度报告者工厂
type ProgressReporterFactory interface {
	Create() ProgressReporter
}

////////////////////////////////////////////////////////////////////////////////

// ProgressReporterHolder 报告者管理器
type ProgressReporterHolder struct {
	factory  ProgressReporterFactory
	reporter ProgressReporter
}

// GetFactory 获取报告者工厂
func (inst *ProgressReporterHolder) GetFactory() ProgressReporterFactory {
	return inst.factory
}

// GetReporter 获取报告者
func (inst *ProgressReporterHolder) GetReporter() (ProgressReporter, error) {
	r := inst.reporter
	if r == nil {
		f := inst.factory
		if f == nil {
			return nil, errors.New("no reporter factory")
		}
		r = f.Create()
		if r == nil {
			return nil, errors.New("no reporter created")
		}
		inst.reporter = r
	}
	return r, nil
}

// SetFactory 设置报告者工厂
func (inst *ProgressReporterHolder) SetFactory(f ProgressReporterFactory) {
	if f != nil {
		inst.factory = f
	}
}

// SetReporter 设置报告者
func (inst *ProgressReporterHolder) SetReporter(r ProgressReporter) {
	inst.reporter = r
}

////////////////////////////////////////////////////////////////////////////////

type DefaultProgressReporterFactory struct {
}

func (inst *DefaultProgressReporterFactory) Create() ProgressReporter {
	return &DefaultProgressReporter{}
}

////////////////////////////////////////////////////////////////////////////////

type DefaultProgressReporter struct {
}

func (inst *DefaultProgressReporter) Report(p *Progress) {

	builder := strings.Builder{}
	builder.WriteString("progress.report")

	builder.WriteString(fmt.Sprint(" name:", p.Name))
	builder.WriteString(fmt.Sprint(" tid:", p.TaskID))
	builder.WriteString(fmt.Sprint(" title:", p.Title))

	builder.WriteString(fmt.Sprint(" unit:", p.Unit))
	builder.WriteString(fmt.Sprint(" min:", p.ValueMin))
	builder.WriteString(fmt.Sprint(" max:", p.ValueMax))
	builder.WriteString(fmt.Sprint(" value:", p.Value))

	vlog.Info(builder.String())
}

func (inst *DefaultProgressReporter) Update(p *Progress) {
	vlog.Info("progress.update state:", p.State, " status:", p.Status)
}

func (inst *DefaultProgressReporter) HandleCancel(f ProgressControlHandlerFn) {

}

func (inst *DefaultProgressReporter) HandlePause(f ProgressControlHandlerFn) {

}

func (inst *DefaultProgressReporter) HandleResume(f ProgressControlHandlerFn) {

}

////////////////////////////////////////////////////////////////////////////////

// GetProgressReporterHolder 取报告者管理器
func GetProgressReporterHolder(ctx context.Context) (*ProgressReporterHolder, error) {
	const key = "github.com/bitwormhole/starter/task/ProgressReporterHolder#binding"
	setter, err := contexts.GetContextSetter(ctx)
	if err != nil {
		return nil, err
	}
	o1 := setter.GetContext().Value(key)
	o2, ok := o1.(*ProgressReporterHolder)
	if ok {
		return o2, nil
	}
	o2 = &ProgressReporterHolder{}
	setter.SetValue(key, o2)
	return o2, nil
}

// GetProgressReporter 取报告者
func GetProgressReporter(ctx context.Context) (ProgressReporter, error) {
	holder, err := GetProgressReporterHolder(ctx)
	if err != nil {
		return nil, err
	}
	return holder.GetReporter()
}
