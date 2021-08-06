package config

import (
	"bytes"
	"embed"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/util"
)

type simpleEmbedResFS struct {
	fs     *embed.FS
	prefix string
}

func (inst *simpleEmbedResFS) computeResPath(path string) string {

	const uriToken = ":/"
	uriTokenIndex := strings.Index(path, uriToken)

	if uriTokenIndex > 0 {
		// 去掉URI.path之前的部分
		uriTokenLength := len(uriToken)
		path = path[uriTokenIndex+uriTokenLength:]
	}

	builder := &util.PathBuilder{}
	builder.EnableRoot(false)
	builder.EnableDoubleDot(false)
	builder.EnableTrim(true)
	builder.AppendPath(inst.prefix)
	builder.AppendPath(path)
	return builder.String()
}

func (inst *simpleEmbedResFS) GetText(path string) (string, error) {
	path = inst.computeResPath(path)
	data, err := inst.fs.ReadFile(path)
	if err != nil {
		return "", err
	}
	text := string(data)
	return text, nil
}

func (inst *simpleEmbedResFS) GetBinary(path string) ([]byte, error) {
	path = inst.computeResPath(path)
	data, err := inst.fs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (inst *simpleEmbedResFS) GetReader(path string) (io.ReadCloser, error) {
	data, err := inst.GetBinary(path)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)
	return ioutil.NopCloser(reader), nil
}

func (inst *simpleEmbedResFS) All() []*collection.Resource {
	return inst.List("/", true)
}

func (inst *simpleEmbedResFS) List(path string, recursive bool) []*collection.Resource {
	// path = inst.computeResPath(path)
	walker := &simpleEmbedTreeWalker{
		fs:        inst.fs,
		basePath:  path,
		recursive: recursive,
		owner:     inst,
	}
	return walker.walk()
}

// CreateEmbedFsResources 从【embed.FS】创建资源组
func CreateEmbedFsResources(fs *embed.FS, pathPrefix string) collection.Resources {
	return &simpleEmbedResFS{
		fs:     fs,
		prefix: pathPrefix,
	}
}

////////////////////////////////////////////////////////////////////////////////

type simpleEmbedTreeWalker struct {
	owner *simpleEmbedResFS

	fs        *embed.FS
	basePath  string
	recursive bool

	results []*collection.Resource
}

func (inst *simpleEmbedTreeWalker) child(name string, parent *collection.Resource) *collection.Resource {
	ch := &collection.Resource{}
	ch.Name = name
	if parent == nil {
		// root
		ch.BasePath = inst.basePath
		ch.AbsolutePath = inst.basePath
		ch.RelativePath = "."
	} else {
		// child
		ch.BasePath = parent.BasePath
		ch.AbsolutePath = parent.AbsolutePath + "/" + name
		ch.RelativePath = parent.RelativePath + "/" + name
	}
	return ch
}

func (inst *simpleEmbedTreeWalker) root() *collection.Resource {
	return inst.child("", nil)
}

func (inst *simpleEmbedTreeWalker) walk() []*collection.Resource {
	node := inst.root()
	inst.results = make([]*collection.Resource, 0)
	inst.walkWithDir(node, 99)
	return inst.results
}

func (inst *simpleEmbedTreeWalker) walkWithDir(node *collection.Resource, limit int) error {
	if limit < 0 {
		return errors.New("path is too deep: " + node.AbsolutePath)
	}
	fullpath := inst.owner.computeResPath(node.AbsolutePath)
	items, err := inst.fs.ReadDir(fullpath)
	if err != nil {
		return err
	}
	for _, item := range items {
		name := item.Name()
		t := item.Type()
		child := inst.child(name, node)
		child.IsDir = item.IsDir()
		if child.IsDir {
			inst.onDir(child)
			if inst.recursive {
				err := inst.walkWithDir(child, limit-1)
				if err != nil {
					return err
				}
			}
		} else if t.IsRegular() {
			inst.onFile(child)
		}
	}

	return nil
}

func (inst *simpleEmbedTreeWalker) onFile(node *collection.Resource) {
	inst.results = append(inst.results, node)
}

func (inst *simpleEmbedTreeWalker) onDir(node *collection.Resource) {
	inst.results = append(inst.results, node)
}

////////////////////////////////////////////////////////////////////////////////
