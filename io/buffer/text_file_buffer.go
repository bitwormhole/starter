package buffer

import (
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

// TextFileBuffer 是一个简单的文本文件读写缓冲区
type TextFileBuffer struct {
	file   fs.Path
	cacheR *textFileBufferCache
}

// Init 初始化缓冲区
func (inst *TextFileBuffer) Init(file fs.Path) {
	inst.file = file
}

// GetText 获取文本
func (inst *TextFileBuffer) GetText(reload bool) string {
	if reload {
		inst.cacheR = nil
	}
	c := inst.getCacheR()
	return c.text
}

// SetText 设置文本
func (inst *TextFileBuffer) SetText(text string, force bool) {

	file := inst.file

	if !force {
		if file.Exists() {
			older := inst.getCacheR()
			if text == older.text {
				return
			}
		}
	}

	newer := &textFileBufferCache{}
	newer.init(file)
	newer.text = text
	err := newer.save()
	if err != nil {
		vlog.Warn(err)
	}
	inst.cacheR = nil
}

func (inst *TextFileBuffer) getCacheR() *textFileBufferCache {
	c := inst.cacheR
	if c != nil {
		if c.isUpToDate() {
			return c
		}
	}
	file := inst.file
	c = &textFileBufferCache{}
	c.init(file)
	err := c.load()
	if err != nil && file.Exists() {
		vlog.Warn(err)
	}
	inst.cacheR = c
	return c
}

////////////////////////////////////////////////////////////////////////////////

type textFileBufferCache struct {
	file fs.Path
	time int64
	size int64
	text string
}

func (inst *textFileBufferCache) init(file fs.Path) {
	inst.file = file
}

func (inst *textFileBufferCache) load() error {
	text, err := inst.file.GetIO().ReadText(nil)
	if err != nil {
		return nil
	}
	inst.text = text
	inst.update()
	return nil
}

func (inst *textFileBufferCache) save() error {
	text := inst.text
	err := inst.file.GetIO().WriteText(text, nil, true)
	if err == nil {
		inst.update()
	}
	return err
}

func (inst *textFileBufferCache) update() {
	file := inst.file
	inst.size = file.Size()
	inst.time = file.LastModTime()
}

func (inst *textFileBufferCache) isUpToDate() bool {
	file := inst.file
	size1 := inst.size
	time1 := inst.time
	size2 := file.Size()
	time2 := file.LastModTime()
	return (size1 == size2) && (time1 == time2)
}
