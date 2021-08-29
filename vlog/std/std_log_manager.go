package std

import "github.com/bitwormhole/starter/vlog"

type stdLogManager struct {
	outputTarget vlog.Channel         // for setter
	outputProxy  vlog.Channel         // for getter
	outputBuffer *managerOutputBuffer // for buffer
}

func (inst *stdLogManager) _Impl() {

}

func (inst *stdLogManager) SetOutput(output vlog.Channel) {

	if output == nil {
		return
	}

	buffer := inst.outputBuffer

	if buffer != nil {
		inst.outputBuffer = nil
		buffer.WriteTo(output)
	}

	inst.outputTarget = output
}

func (inst *stdLogManager) GetOutput() vlog.Channel {
	out := inst.outputProxy
	if out == nil {
		out = (&managerOutputProxy{}).init(inst)
		inst.outputProxy = out
	}
	return out
}

////////////////////////////////////////////////////////////////////////////////

type managerOutputProxy struct {
	manager *stdLogManager
}

func (inst *managerOutputProxy) init(manager *stdLogManager) vlog.Channel {

	inst.manager = manager

	buffer := &managerOutputBuffer{}
	buffer.init()
	manager.outputBuffer = buffer
	manager.outputTarget = buffer

	return inst
}

func (inst *managerOutputProxy) IsLevelEnabled(l vlog.Level) bool {
	return inst.manager.outputTarget.IsLevelEnabled(l)
}

func (inst *managerOutputProxy) Write(r *vlog.Record) {
	inst.manager.outputTarget.Write(r)
}

////////////////////////////////////////////////////////////////////////////////

type managerOutputBuffer struct {
	list []*vlog.Record
}

func (inst *managerOutputBuffer) init() vlog.Channel {
	return inst
}

func (inst *managerOutputBuffer) IsLevelEnabled(l vlog.Level) bool {
	return true
}

func (inst *managerOutputBuffer) Write(r *vlog.Record) {
	inst.list = append(inst.list, r)
}

func (inst *managerOutputBuffer) WriteTo(target vlog.Channel) {
	if target == nil {
		return
	}
	src := inst.list
	for _, item := range src {
		if !target.IsLevelEnabled(item.Level) {
			continue
		}
		target.Write(item)
	}
}

////////////////////////////////////////////////////////////////////////////////

var theStdLogManagerInstance *stdLogManager

func getManager() *stdLogManager {
	inst := theStdLogManagerInstance
	if inst == nil {
		inst = &stdLogManager{}
		theStdLogManagerInstance = inst
	}
	return inst
}
