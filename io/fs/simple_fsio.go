package fs

import (
	"io/ioutil"
	"os"
)

type innerFileIO struct {
	path *innerPath
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

func (inst *innerFileIO) WriteBinary(data []byte, mode IoMode) error {
	var perm os.FileMode = 0755
	if mode != nil {
		perm = mode.Perm()
	}
	filename := inst.path.path
	return ioutil.WriteFile(filename, data, perm)
}

func (inst *innerFileIO) WriteText(text string, mode IoMode) error {
	data := []byte(text)
	return inst.WriteBinary(data, mode)
}
