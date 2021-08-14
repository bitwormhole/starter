package tests

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

// TestDirectoryLoader 测试目录加载器
type TestDirectoryLoader struct {
	r collection.Resources
	t *testing.T
}

// Init 初始化
func (inst *TestDirectoryLoader) Init(r collection.Resources, t *testing.T) {
	inst.r = r
	inst.t = t
}

// LoadFromFolder 从资源组中的文件夹加载测试目录
func (inst *TestDirectoryLoader) LoadFromFolder(path string) (fs.Path, error) {

	t := inst.t
	r := inst.r
	temp := fs.Default().GetPath(t.TempDir())
	items := r.List(path, true)
	target := temp.GetChild("./" + path)

	for _, item := range items {
		dst := target.GetChild("./" + item.RelativePath)
		if item.IsDir {
			dst.Mkdirs()
			continue
		} else {
			dir := dst.Parent()
			if !dir.Exists() {
				dir.Mkdirs()
			}
		}
		data, err := r.GetBinary(item.AbsolutePath)
		if err != nil {
			return nil, err
		}
		err = dst.GetIO().WriteBinary(data, nil)
		if err != nil {
			return nil, err
		}
	}

	return target, nil
}

// LoadFromZipFile 从资源组中的压缩文件加载测试目录
func (inst *TestDirectoryLoader) LoadFromZipFile(zipfile string) (fs.Path, error) {
	r := inst.r
	temp := fs.Default().GetPath(inst.t.TempDir())
	data, err := r.GetBinary(zipfile)
	if err != nil {
		return nil, err
	}
	target := temp.GetChild("./" + zipfile)
	reader := inst.createReaderAt(data)
	size := len(data)
	zipReader, err := zip.NewReader(reader, int64(size))
	if err != nil {
		return nil, err
	}
	items := zipReader.File
	for _, item := range items {
		name := item.Name
		output := target.GetChild(name)
		if item.Mode().IsDir() {
			output.Mkdirs()
			continue
		}
		src, err := item.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()
		entity := inst.readAll(src)
		err = output.GetIO().WriteBinary(entity, nil)
		if err != nil {
			return nil, err
		}
		vlog.Debug("write to ", output.Path())
		vlog.Debug("    size=", item.FileInfo().Size())
	}
	return target, nil
}

func (inst *TestDirectoryLoader) readAll(r io.ReadCloser) []byte {
	dst := &bytes.Buffer{}
	buffer := make([]byte, 1024)
	for {
		cb, err := r.Read(buffer)
		if cb > 0 {
			dst.Write(buffer[0:cb])
		}
		if err != nil {
			break
		}
	}
	return dst.Bytes()
}

func (inst *TestDirectoryLoader) createReaderAt(data []byte) io.ReaderAt {
	r := &zipReaderAt{}
	return r.init(data)
}

////////////////////////////////////////////////////////////////////////////////

type zipReaderAt struct {
	data   []byte
	length int64
}

func (inst *zipReaderAt) init(data []byte) io.ReaderAt {
	inst.data = data
	inst.length = int64(len(data))
	return inst
}

func (inst *zipReaderAt) ReadAt(b []byte, p int64) (int, error) {
	wantSize := int64(len(b))
	if 0 <= p && p < (inst.length) {
		size := inst.length - p
		if size > wantSize {
			size = wantSize
		}
		for i := int64(0); i < size; i++ {
			b[i] = inst.data[p+i]
		}
		return int(size), nil
	}
	return 0, errors.New("out of buffer")
}
