package configen

import "github.com/bitwormhole/starter/application"

type Runner struct {
}

func (inst *Runner) Run(context application.RuntimeContext) error {
	return Main(context)
}
