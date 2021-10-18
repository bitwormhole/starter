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

func (inst *innerWindowsPlatform) isAbsolute(path string) bool {

	const headLen = 3
	if len(path) < headLen {
		return false
	}
	head := path[0:headLen]
	array := []byte(head)

	c0 := array[0]
	c1 := array[1]
	c2 := array[2]

	ok0 := (('a' <= c0 && c0 <= 'z') || ('A' <= c0 && c0 <= 'Z'))
	ok1 := c1 == ':'
	ok2 := (c2 == '\\') || (c2 == '/')

	return ok0 && ok1 && ok2
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
