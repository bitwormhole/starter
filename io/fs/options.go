package fs

import "os"

// Options 包含文件系统IO的选项
type Options struct {
	Flag int
	Mode os.FileMode
}

// Clone 返回一个Options的副本
func (inst *Options) Clone() *Options {
	op2 := &Options{}
	op2.Flag = inst.Flag
	op2.Mode = inst.Mode
	return op2
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
