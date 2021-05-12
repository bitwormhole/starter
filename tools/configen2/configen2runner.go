package configen2

import (
	"errors"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
)

type Runner struct {
}

func (inst *Runner) Run(context application.RuntimeContext) error {

	pwd, err := context.GetEnvironment().GetEnv("PWD")
	if err != nil {
		return err
	}
	pwd_path := fs.Default().GetPath(pwd)

	ctx := &configen2context{pwd: pwd_path}
	proc := &configen2process{context: ctx}
	return proc.run()
}

func (inst *Runner) RunWithPWD(pwd string) error {
	if pwd == "" {
		return errors.New("PWD==nil")
	}
	pwd_path := fs.Default().GetPath(pwd)
	ctx := &configen2context{pwd: pwd_path}
	proc := &configen2process{context: ctx}
	return proc.run()
}
