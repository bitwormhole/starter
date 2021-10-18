package buffers

import (
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

// TextFileBuffer 是一个简单的文本文件读写缓冲区
type TextFileBuffer struct {
	file  fs.Path
	cache *textFileBufferCache
}

// Init 初始化缓冲区
func (inst *TextFileBuffer) Init(file fs.Path) {
	inst.file = file
	inst.cache = nil
}

// GetText 获取文本
func (inst *TextFileBuffer) GetText(reload bool) string {
	if reload {
		inst.cache = nil
	}
	c := inst.getCache()
	return c.text
}

// SetText 设置文本
func (inst *TextFileBuffer) SetText(text string, force bool) {

	file := inst.file

	if !force {
		cache := inst.getCache()
		if text == cache.text && cache.exists {
			return
		}
	}

	err := file.GetIO().WriteText(text, nil, true)
	if err != nil {
		vlog.Warn(err)
	}
	inst.cache = nil
}

func (inst *TextFileBuffer) getCache() *textFileBufferCache {
	c := inst.cache
	if c != nil {
		if c.isUpToDate() {
			return c
		}
	}
	file := inst.file
	c = &textFileBufferCache{}
	c.init(file)
	err := c.load()
	if err != nil && c.exists {
		vlog.Warn(err)
	}
	inst.cache = c
	return c
}

////////////////////////////////////////////////////////////////////////////////

type textFileBufferCache struct {
	file   fs.Path
	time   int64
	size   int64
	text   string
	exists bool
}

func (inst *textFileBufferCache) init(file fs.Path) {
	inst.file = file
}

func (inst *textFileBufferCache) load() error {
	text, err := inst.file.GetIO().ReadText(nil)
	if err != nil {
		inst.exists = false
		return nil
	}
	inst.text = text
	inst.update()
	return nil
}

func (inst *textFileBufferCache) update() {
	file := inst.file
	inst.size = file.Size()
	inst.time = file.LastModTime()
	inst.exists = file.Exists()
}

func (inst *textFileBufferCache) isUpToDate() bool {
	file := inst.file
	if !file.Exists() {
		return false
	}
	size1 := inst.size
	time1 := inst.time
	size2 := file.Size()
	time2 := file.LastModTime()
	return (size1 == size2) && (time1 == time2)
}
