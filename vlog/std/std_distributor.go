package std

import "github.com/bitwormhole/starter/vlog"

// Distributor 把 rec 分发到多个channel
type Distributor struct {
	Channels []vlog.Channel
}

func (inst *Distributor) _Impl() vlog.Writer {
	return inst
}

func (inst *Distributor) Write(rec *vlog.Record) {
	chls := inst.Channels
	for _, ch := range chls {
		if ch == nil {
			continue
		}
		ch.Write(rec)
	}
}
