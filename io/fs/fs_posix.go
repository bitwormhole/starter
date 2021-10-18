package fs

import (
	"strings"

	"github.com/bitwormhole/starter/util"
)

type innerPosixPlatform struct{}

func (inst *innerPosixPlatform) Roots() []string {
	list := make([]string, 1)
	list[0] = "/"
	return list
}

func (inst *innerPosixPlatform) isAbsolute(path string) bool {
	return strings.HasPrefix(path, "/")
}

func (inst *innerPosixPlatform) PathSeparatorChar() rune {
	return ':'
}

func (inst *innerPosixPlatform) SeparatorChar() rune {
	return '/'
}

func (inst *innerPosixPlatform) normalizePath(path string) (string, error) {
	sep := string(inst.SeparatorChar())
	pb := &util.PathBuilder{}
	pb.SetSeparator(sep)
	pb.AppendPath(path)
	text, err := pb.Create("", "")
	if err != nil {
		return "", err
	}
	return sep + text, nil
}
