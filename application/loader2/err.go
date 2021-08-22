package loader2

import (
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/vlog"
)

type defaultErrorHandler struct{}

func (inst *defaultErrorHandler) _Impl() lang.ErrorHandler {
	return inst
}

func (inst *defaultErrorHandler) OnError(err error) {
	vlog.Error(err.Error())
}
