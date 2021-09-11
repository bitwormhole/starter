package fs

import (
	"io"
	"io/ioutil"
	"os"
)

type innerFileIO struct {
	path *innerPath
}

func (inst *innerFileIO) _Impl() FileIO {
	return inst
}

func (inst *innerFileIO) prepareOptionsForRead(opt *Options) *Options {
	if opt != nil {
		return opt
	}
	return inst.path.FileSystem().DefaultReadOptions()
}

func (inst *innerFileIO) prepareOptionsForWrite(opt *Options) *Options {
	if opt != nil {
		return opt
	}
	return inst.path.FileSystem().DefaultWriteOptions()
}

func (inst *innerFileIO) tryMkdirs(mkdirs bool) {
	if mkdirs {
		dir := inst.Path().Parent()
		if !dir.Exists() {
			dir.Mkdirs()
		}
	}
}

func (inst *innerFileIO) Path() Path {
	return inst.path
}

func (inst *innerFileIO) ReadBinary(opt *Options) ([]byte, error) {
	opt = inst.prepareOptionsForRead(opt)
	filename := inst.path.path
	return ioutil.ReadFile(filename)
}

func (inst *innerFileIO) ReadText(opt *Options) (string, error) {
	opt = inst.prepareOptionsForRead(opt)
	filename := inst.path.path
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (inst *innerFileIO) WriteBinary(data []byte, opt *Options, mkdirs bool) error {
	inst.tryMkdirs(mkdirs)
	opt = inst.prepareOptionsForWrite(opt)
	filename := inst.path.path
	return ioutil.WriteFile(filename, data, opt.Mode)
}

func (inst *innerFileIO) WriteText(text string, opt *Options, mkdirs bool) error {
	// inst.tryMkdirs(mkdirs)
	opt = inst.prepareOptionsForWrite(opt)
	data := []byte(text)
	return inst.WriteBinary(data, opt, mkdirs)
}

func (inst *innerFileIO) OpenReader(opt *Options) (io.ReadCloser, error) {
	opt = inst.prepareOptionsForRead(opt)
	path := inst.Path()
	f, err := os.OpenFile(path.Path(), opt.Flag, opt.Mode)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (inst *innerFileIO) OpenWriter(opt *Options, mkdirs bool) (io.WriteCloser, error) {
	inst.tryMkdirs(mkdirs)
	opt = inst.prepareOptionsForWrite(opt)
	path := inst.Path()
	if opt.Create {
		return os.Create(path.Path())
	}
	f, err := os.OpenFile(path.Path(), opt.Flag, opt.Mode)
	if err != nil {
		return nil, err
	}
	return f, nil
}
