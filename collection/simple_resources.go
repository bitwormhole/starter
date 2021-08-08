package collection

import (
	"errors"
	"io"
	"strings"

	"github.com/bitwormhole/starter/util"
)

////////////////////////////////////////////////////////////////////////////////

type simpleResources struct {
	table map[string]Res
}

func (inst *simpleResources) GetText(name string) (string, error) {
	r, err := inst.Get(name)
	if err != nil {
		return "", err
	}
	return r.ReadText()
}

func (inst *simpleResources) GetBinary(name string) ([]byte, error) {
	r, err := inst.Get(name)
	if err != nil {
		return nil, err
	}
	return r.ReadBinary()
}

func (inst *simpleResources) GetReader(name string) (io.ReadCloser, error) {
	r, err := inst.Get(name)
	if err != nil {
		return nil, err
	}
	return r.Reader()
}

func (inst *simpleResources) Get(path string) (Res, error) {
	path = inst.normalizePath(path)
	table := inst.table
	if table != nil {
		res := table[path]
		if res != nil {
			return res, nil
		}
	}
	return nil, errors.New("no resource, path=" + path)
}

func (inst *simpleResources) Clear() {
	inst.table = make(map[string]Res)
}

func (inst *simpleResources) Export(table map[string]Res) map[string]Res {
	src := inst.table
	dst := table
	if dst == nil {
		dst = make(map[string]Res)
	}
	if src == nil {
		return dst
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (inst *simpleResources) Import(src map[string]Res, override bool) {
	dst := inst.table
	if dst == nil {
		dst = make(map[string]Res)
		inst.table = dst
	}
	if src == nil {
		return
	}
	for k, v := range src {
		path := inst.normalizePath(k)
		older := dst[path]
		if older == nil {
			dst[path] = v
		} else if override {
			dst[path] = v
		}
	}
}

// 列出所有资源的路径, 相当于{{ List("/",true) }}
func (inst *simpleResources) All() []*Resource {
	return inst.List("/", true)
}

// 列出所有资源的路径
func (inst *simpleResources) List(path string, recursive bool) []*Resource {
	src := inst.table
	dst := make([]*Resource, 0)
	if src == nil {
		return dst
	}
	prefix := inst.normalizePath(path)
	if len(prefix) > 0 && !strings.HasSuffix(prefix, "/") {
		prefix = prefix + "/"
	}
	for location, res := range src {
		if !strings.HasPrefix(location, prefix) {
			continue
		}
		info := inst.makeResInfo(location, res, prefix)
		if recursive {
			dst = append(dst, info)
		} else if !inst.isDeepNode(info) {
			dst = append(dst, info)
		}
	}
	return dst
}

func (inst *simpleResources) isDeepNode(r *Resource) bool {
	path := r.RelativePath
	return strings.Contains(path, "/")
}

func (inst *simpleResources) makeResInfo(location string, src Res, prefix string) *Resource {
	name := location
	index := strings.LastIndex(location, "/")
	if index >= 0 {
		name = location[index+1:]
	}
	dst := &Resource{}
	dst.IsDir = src.IsDir()
	dst.RelativePath = location[len(prefix):]
	dst.AbsolutePath = location
	dst.BasePath = prefix
	dst.Name = name
	return dst
}

func (inst *simpleResources) normalizePath(path string) string {
	const token = ":/"
	index := strings.Index(path, token)
	if index >= 0 {
		path = path[index+len(token):]
	}
	builder := &util.PathBuilder{}
	builder.AppendPath(path)
	return builder.String()
}

func (inst *simpleResources) init() Resources {
	inst.Clear()
	return inst
}

////////////////////////////////////////////////////////////////////////////////

// CreateResources 创建一个空的资源组
func CreateResources() Resources {
	inst := &simpleResources{}
	return inst.init()
}
