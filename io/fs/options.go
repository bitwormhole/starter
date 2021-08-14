package fs

import "os"

// Options 包含文件系统IO的选项
type Options struct {
	Flag int
	Mode os.FileMode
}

// Normalize 返回一个正常的Options
func (inst *Options) Normalize() *Options {

	if inst != nil {
		return inst
	}

	opt := &Options{}

	// todo : opt.Mode = io.

	return opt
}
