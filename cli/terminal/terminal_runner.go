package terminal

import (
	"bytes"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/cli"
)

// Run 运行一个内嵌的terminal循环，直到exit
func Run(app application.Context) error {

	cli.SetupConsole(app, nil)

	// app.GetComponents().FindComponentsWithFilter("", func(name string, holder application.ComponentHolder) bool {
	// 	return true
	// })

	console, err := cli.GetConsole(app)
	if err != nil {
		return err
	}

	ctx := &Context{}
	ctx.app = app
	ctx.console = console
	ctx.ctx = app
	ctx.prompt = "$"
	ctx.exit = false
	ctx.client = nil // todo

	runner := &runner{context: ctx}

	return runner.run()
}

////////////////////////////////////////////////////////////////////////////////

type runner struct {
	context *Context
}

func (inst *runner) run() error {

	ctx := inst.context
	in := ctx.console.Input()
	buffer := make([]byte, 1)
	line := bytes.Buffer{}

	for {
		n, err := in.Read(buffer)
		if err != nil {
			return err
		}
		if n < 1 {
			break
		}
		b := buffer[0]
		if b == '\n' || b == '\r' {
			str := line.String()
			line.Reset()
			err = inst.handleLine(str)
			if err != nil {
				return err
			}
		} else {
			line.Write(buffer)
		}
		if ctx.exit {
			break
		}
	}

	return nil
}

func (inst *runner) handleLine(line string) error {
	client := inst.context.client
	return client.ExecuteScript(line)
}
