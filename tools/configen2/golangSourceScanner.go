package configen2

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/io/fs"
)

type golangSourceScanner struct {
	context *configen2context
}

func (inst *golangSourceScanner) scan() error {
	dir := inst.context.inputFile.Parent()
	return inst.scanDir(dir)
}

func (inst *golangSourceScanner) scanDir(dir fs.Path) error {
	namelist := dir.ListNames()
	for index := range namelist {
		name := namelist[index]
		if strings.HasSuffix(name, ".go") {
			err := inst.scanSourceFile(dir.GetChild(name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (inst *golangSourceScanner) scanSourceFile(file fs.Path) error {

	fmt.Println("scan golang source", file.Path())

	loader := &comLoader{}
	com, err := loader.loadFromFile(file)
	if err != nil {
		return err
	}

	inst.context.com = com
	return nil
}
