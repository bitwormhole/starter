package config

import (
	"embed"
	"errors"
	"io"
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
	uriTokenLength := len(uriToken)

	if uriTokenIndex > 0 {
		// 去掉URI.path之前的部分
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
	return nil, nil
}

func (inst *simpleEmbedResFS) All() []string {
	walker := &simpleEmbedTreeWalker{
		fs:     inst.fs,
		prefix: inst.prefix,
	}
	return walker.walk()
}

func CreateEmbedFsResources(fs *embed.FS, pathPrefix string) collection.Resources {
	return &simpleEmbedResFS{
		fs:     fs,
		prefix: pathPrefix,
	}
}

////////////////////////////////////////////////////////////////////////////////

type simpleEmbedTreeWalker struct {
	fs      *embed.FS
	prefix  string
	results []string
}

func (inst *simpleEmbedTreeWalker) walk() []string {
	inst.results = make([]string, 0)
	inst.walkWithDir(inst.prefix, 99)
	return inst.results
}

func (inst *simpleEmbedTreeWalker) walkWithDir(path string, limit int) error {
	if limit < 0 {
		return errors.New("path is too deep: " + path)
	}
	items, err := inst.fs.ReadDir(path)
	if err != nil {
		return err
	}
	for index := range items {
		item := items[index]
		name := item.Name()
		t := item.Type()
		if item.IsDir() {
			err := inst.walkWithDir(path+"/"+name, limit-1)
			if err != nil {
				return err
			}
		} else if t.IsRegular() {
			inst.onFile(path + "/" + name)
		}
	}
	return nil
}

func (inst *simpleEmbedTreeWalker) onFile(path string) {
	len1 := len(inst.prefix)
	path = path[len1:]
	inst.results = append(inst.results, "res://"+path)
	// fmt.Println("onFile: " + path)
}

////////////////////////////////////////////////////////////////////////////////
