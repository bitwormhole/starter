package config

import (
	"bytes"
	"embed"
	"errors"
	"io"
	"io/ioutil"
	"log"

	"github.com/bitwormhole/starter/collection"
)

////////////////////////////////////////////////////////////////////////////////

// LoadResourcesFromEmbedFS 从嵌入的FS加载资源组
func LoadResourcesFromEmbedFS(fs *embed.FS, path string) collection.Resources {
	loader := &simpleEmbedTreeLoader{}
	all, err := loader.load(fs, path)
	if err != nil {
		log.Println("[WARN] LoadResourcesFromEmbedFS error: " + err.Error())
		return collection.CreateResources()
	}
	return all
}

////////////////////////////////////////////////////////////////////////////////

type simpleEmbedRes struct {
	fs    *embed.FS
	src   string
	isdir bool

	data []byte // cache
}

func (inst *simpleEmbedRes) _Impl() collection.Res {
	return inst
}

func (inst *simpleEmbedRes) Length() int64 {
	data, err := inst.getbin()
	if err != nil {
		return 0
	}
	size := len(data)
	return int64(size)
}

func (inst *simpleEmbedRes) Exists() bool {
	return true
}

func (inst *simpleEmbedRes) IsDir() bool {
	return inst.isdir
}

func (inst *simpleEmbedRes) IsFile() bool {
	return !inst.isdir
}

func (inst *simpleEmbedRes) getbin() ([]byte, error) {
	if inst.isdir {
		return nil, errors.New("node is a dir, src=" + inst.src)
	}
	data := inst.data
	if data == nil {
		bin, err := inst.fs.ReadFile(inst.src)
		if err != nil {
			return nil, err
		}
		data = bin
		inst.data = bin
	}
	return data, nil
}

func (inst *simpleEmbedRes) ReadBinary() ([]byte, error) {
	data, err := inst.getbin()
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(data))
	copy(dst, data)
	return dst, nil
}

func (inst *simpleEmbedRes) ReadText() (string, error) {
	data, err := inst.getbin()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (inst *simpleEmbedRes) Reader() (io.ReadCloser, error) {
	data, err := inst.getbin()
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)
	return ioutil.NopCloser(reader), nil
}

////////////////////////////////////////////////////////////////////////////////

type simpleEmbedTreeLoader struct {
}

func (inst *simpleEmbedTreeLoader) load(fs *embed.FS, path string) (collection.Resources, error) {

	table := make(map[string]collection.Res)
	result := collection.CreateResources()
	const depthLimit = 99

	err := inst.walkDir(fs, path, "/", depthLimit, table)
	if err != nil {
		return nil, err
	}

	result.Import(table, false)
	return result, nil
}

func (inst *simpleEmbedTreeLoader) walkDir(fs *embed.FS, pathFs string, pathRes string, depthLimit int, dst map[string]collection.Res) error {
	if depthLimit < 1 {
		return errors.New("path is too deep, path=" + pathFs)
	}
	items, err := fs.ReadDir(pathFs)
	if err != nil {
		return err
	}
	for _, item := range items {
		name := item.Name()
		isfile := item.Type().IsRegular()
		if item.IsDir() {
			err := inst.walkDir(fs, pathFs+"/"+name, pathRes+"/"+name, depthLimit-1, dst)
			if err != nil {
				return err
			}
		} else if isfile {
			inst.onFile(fs, pathFs+"/"+name, pathRes+"/"+name, dst)
		}
	}
	return nil
}

func (inst *simpleEmbedTreeLoader) onFile(fs *embed.FS, pathFs string, pathRes string, dst map[string]collection.Res) {
	item := &simpleEmbedRes{}
	item.src = pathFs
	item.fs = fs
	item.isdir = false
	dst[pathRes] = item
}

////////////////////////////////////////////////////////////////////////////////
