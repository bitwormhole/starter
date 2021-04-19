package fs

import (
	"os"
)

type innerFileMeta struct {
	exists bool
	mode   os.FileMode
	info   os.FileInfo
}

func (inst *innerFileMeta) load(path string) {
	info, err := os.Stat(path)
	if err == nil {
		inst.exists = true
		inst.mode = info.Mode()
		inst.info = info
	} else {
		inst.exists = os.IsExist(err)
		inst.mode = 0
		inst.info = nil
	}
}

func (inst *innerFileMeta) Exists() bool {
	return inst.exists
}

func (inst *innerFileMeta) LastModTime() int64 {
	info := inst.info
	if info == nil {
		return 0
	}
	sec := info.ModTime().Unix()
	return sec * 1000
}

func (inst *innerFileMeta) Size() int64 {
	info := inst.info
	if info == nil {
		return 0
	}
	return info.Size()
}

func (inst *innerFileMeta) IsFile() bool {
	info := inst.info
	if info == nil {
		return false
	}
	return !info.IsDir()
}

func (inst *innerFileMeta) IsDir() bool {
	info := inst.info
	if info == nil {
		return false
	}
	return info.IsDir()
}

func (inst *innerFileMeta) IsSymlink() bool {
	if inst.exists {
		return (inst.mode & os.ModeSymlink) != 0
	}
	return false
}

func (inst *innerFileMeta) CanRead() bool {

	return false
}

func (inst *innerFileMeta) CanWrite() bool {

	return false
}

func (inst *innerFileMeta) CanExecute() bool {
	return false
}
