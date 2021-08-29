package std

import (
	"fmt"
	"strings"
	"time"

	"github.com/bitwormhole/starter/vlog"
)

type DefaultFormatter struct {
}

func (inst *DefaultFormatter) _Impl() vlog.Formatter {
	return inst
}

func (inst *DefaultFormatter) Format(r *vlog.Record) string {

	tt := time.Unix(r.Timestamp/1000, 0)
	args := r.Arguments

	builder := strings.Builder{}
	builder.WriteString(tt.String())
	builder.WriteString(" [" + r.Level.String() + "] ")
	builder.WriteString(fmt.Sprint(args...))
	return builder.String()
}
