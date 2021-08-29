package std

import "github.com/bitwormhole/starter/vlog"

// Context 日志上下文
type Context interface {
	GetDefaultLevel() vlog.Level
	GetDefaultFormatter() vlog.Formatter
	GetMainChannel() vlog.Channel
}

// DefaultContext 日志上下文
type DefaultContext struct {
	DefaultLevel     string
	DefaultFormatter vlog.Formatter
	Channels         []vlog.Channel
	MainChannel      vlog.Channel

	defaultLevel vlog.Level
}

// LogChannel 日志通道
type LogChannel struct {

	//  public
	Context   Context
	Name      string
	Enable    bool
	Filters   []vlog.Filter
	Writer    vlog.Writer
	Level     string
	Formatter vlog.Formatter

	// private
	myLevel vlog.Level
	chain   vlog.FilterChain
}

////////////////////////////////////////////////////////////////////////////////

func (inst *DefaultContext) _Impl() Context {
	return inst
}

// GetDefaultLevel 取默认的日志等级
func (inst *DefaultContext) GetDefaultLevel() vlog.Level {
	return inst.defaultLevel
}

// GetDefaultFormatter 取默认的 formatter
func (inst *DefaultContext) GetDefaultFormatter() vlog.Formatter {
	return inst.DefaultFormatter
}

// GetMainChannel 取主通道
func (inst *DefaultContext) GetMainChannel() vlog.Channel {
	return inst.MainChannel
}

////////////////////////////////////////////////////////////////////////////////

func (inst *LogChannel) _Impl() vlog.Channel {
	return inst
}

func (inst *LogChannel) Write(r *vlog.Record) {
	if r == nil {
		return
	}
	if !inst.IsLevelEnabled(r.Level) {
		return
	}
	next := inst.Writer
	if next == nil {
		return
	}
	next.Write(r)
}

func (inst *LogChannel) getSettingLevel() vlog.Level {
	value := inst.myLevel
	if value == 0 {
		str := inst.Level
		v2, err := vlog.ParseLevel(str)
		if err == nil {
			value = v2
		} else {
			value = inst.Context.GetDefaultLevel()
		}
		inst.myLevel = value
	}
	return value
}

// IsLevelEnabled 判断给定的日志等级是否可用
func (inst *LogChannel) IsLevelEnabled(level vlog.Level) bool {
	if !inst.Enable {
		return false
	}
	setting := inst.getSettingLevel()
	return (setting <= level)
}
