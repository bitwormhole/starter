package elements

import (
	"fmt"

	"github.com/bitwormhole/starter/application"
)

////////////////////////////////////////////////////////////////////////////////

type ExitCoder1 struct {
}

func (inst *ExitCoder1) _impl_exit_coder() application.ExitCodeGenerator {
	return inst
}

func (inst *ExitCoder1) ExitCode() int {
	code := 11
	fmt.Println("ExitCoder1.exit()", code)
	return code
}

func (inst *ExitCoder1) Priority() int {
	return 11
}

////////////////////////////////////////////////////////////////////////////////

type ExitCoder2 struct {
}

func (inst *ExitCoder2) _impl_exit_coder() application.ExitCodeGenerator {
	return inst
}

func (inst *ExitCoder2) ExitCode() int {
	code := 22
	fmt.Println("ExitCoder2.exit()", code)
	return code
}

func (inst *ExitCoder2) Priority() int {
	return 22
}
