package cli

import (
	"strings"
	"testing"

	"github.com/bitwormhole/starter/vlog"
)

func TestTaskListBuilder(t *testing.T) {

	sb := strings.Builder{}

	sb.WriteString("\n")
	sb.WriteString("   \t  example   \n")
	sb.WriteString("   example  foo bar   \n")
	sb.WriteString("   example -a -b -c  --x --y --z  \n")
	sb.WriteString("\n")
	sb.WriteString("    example  -m='hello,world'    \n")
	sb.WriteString("\n")

	builder := &taskListBuilder{}
	builder.parseScript(sb.String())

	list := builder.create()
	for _, unit := range list {
		args := unit.Arguments
		vlog.Info("command-line num:", unit.LineNumber, " text:", unit.CommandLine)
		if args == nil {
			continue
		}
		for j, arg := range args {
			vlog.Info("      args[", j, "] = ", arg)
		}
	}
}
