package configen

import (
	"errors"
	"os"

	"github.com/bitwormhole/starter/io/fs"
)

func Main(args []string) error {

	pwd := os.Getenv("PWD")
	if pwd == "" {
		return errors.New("no env:'PWD'")
	}
	context := &configenContext{}
	context.inputFileName = "starter.config"
	context.pwd = fs.Default().GetPath(pwd)

	proc := &configenProcess{context: context}
	return proc.run()
}
