package fs

import (
	"errors"
	"os"
	"strings"

	"github.com/bitwormhole/starter/util"
)

type innerWindowsPlatform struct{}

func (inst *innerWindowsPlatform) isRootExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func (inst *innerWindowsPlatform) normalizePath(path string) (string, error) {
	sep := string(inst.SeparatorChar())
	pb := &util.PathBuilder{}
	pb.SetSeparator(sep)
	pb.AppendPath(path)
	text, err := pb.Create("", "")
	if err != nil {
		return "", err
	}
	if text == "" {
		return "", errors.New("path==[empty]")
	}
	if !strings.Contains(text, sep) {
		return text + sep, nil
	}
	return text, nil
}

func (inst *innerWindowsPlatform) Roots() []string {

	const driveA rune = 'A'
	const driveZ rune = 'Z'
	list := make([]string, 0)

	for drive := driveA; drive <= driveZ; drive++ {
		path := string(drive) + ":\\"
		if inst.isRootExists(path) {
			list = append(list, path)
		}
	}

	return list
}

func (inst *innerWindowsPlatform) PathSeparatorChar() rune {
	return ';'
}

func (inst *innerWindowsPlatform) SeparatorChar() rune {
	return '\\'
}
