package fs

type innerPosixPlatform struct{}

func (inst *innerPosixPlatform) Roots() []string {
	list := make([]string, 1)
	list[0] = "/"
	return list
}

func (inst *innerPosixPlatform) PathSeparatorChar() rune {
	return ':'
}

func (inst *innerPosixPlatform) SeparatorChar() rune {
	return '/'
}
