package configen

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
)

func Main(ctx application.RuntimeContext) error {

	pwd, err := ctx.GetEnvironment().GetEnv("PWD")
	if err != nil {
		return err
	}

	context := &configenContext{}
	context.inputFileName = "starter.config"
	context.pwd = fs.Default().GetPath(pwd)

	proc := &configenProcess{context: context}
	return proc.run()
}
