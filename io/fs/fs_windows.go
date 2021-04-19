package fs

import "os"

type innerWindowsPlatform struct{}

func (inst *innerWindowsPlatform) isRootExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
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
