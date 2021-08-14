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

func (inst *innerFileIO) Path() Path {
	return inst.path
}

func (inst *innerFileIO) ReadBinary() ([]byte, error) {
	filename := inst.path.path
	return ioutil.ReadFile(filename)
}

func (inst *innerFileIO) ReadText() (string, error) {
	filename := inst.path.path
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (inst *innerFileIO) WriteBinary(data []byte, opt *Options) error {
	opt = opt.Normalize()
	filename := inst.path.path
	return ioutil.WriteFile(filename, data, opt.Mode)
}

func (inst *innerFileIO) WriteText(text string, opt *Options) error {
	data := []byte(text)
	return inst.WriteBinary(data, opt)
}

func (inst *innerFileIO) OpenReader() (io.ReadCloser, error) {
	path := inst.Path()
	f, err := os.OpenFile(path.Path(), 0, 0)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (inst *innerFileIO) OpenWriter(opt *Options, mkdirs bool) (io.WriteCloser, error) {
	path := inst.Path()
	if mkdirs {
		dir := path.Parent()
		if dir.Exists() {
			dir.Mkdirs()
		}
	}
	f, err := os.OpenFile(path.Path(), 0, 0)
	if err != nil {
		return nil, err
	}
	return f, nil
}
