package tests

import (
	"archive/zip"
	"crypto/sha256"
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
)

type testingDataLoader interface {
	Load(ctx TestingRuntime, name string, dst fs.Path) (fs.Path, error)
}

////////////////////////////////////////////////////////////////////////////////

type defaultTestingDataLoader struct {
}

func (inst *defaultTestingDataLoader) prepareDestinationDir(rt TestingRuntime, name string, dst fs.Path) fs.Path {

	if dst != nil {
		return dst
	}

	sum := sha256.Sum256([]byte(name))
	hash := util.StringifyBytes(sum[0:3])

	temp := rt.T().TempDir()
	ref := fs.Default().GetPath(temp + "/" + name)
	parent := ref.Parent()
	shortName := ref.Name()

	return parent.GetChild(shortName + "-" + hash)
}

func (inst *defaultTestingDataLoader) Load(ctx TestingRuntime, name string, dst fs.Path) (fs.Path, error) {

	if name == "" {
		// disable testing data
		return nil, nil
	}

	dst = inst.prepareDestinationDir(ctx, name, dst)
	if strings.HasSuffix(name, ".zip") {
		loader := &zipTestingDirLoader{}
		return loader.Load(ctx, name, dst)
	}

	loader := &sparseTestingDirLoader{}
	return loader.Load(ctx, name, dst)
}

////////////////////////////////////////////////////////////////////////////////

type zipTestingDirLoading struct {
	sourceZipName string
	resources     collection.Resources
	tmpZipFile    fs.Path
	targetDir     fs.Path
}

type zipTestingDirLoader struct {
}

func (inst *zipTestingDirLoader) Load(rt TestingRuntime, name string, dst fs.Path) (fs.Path, error) {

	res := rt.Context().GetResources()
	parent := dst.Parent()
	simpleName := dst.Name()
	tmpZip := parent.GetChild(simpleName + ".zip")

	loading := &zipTestingDirLoading{
		resources:     res,
		sourceZipName: name,
		targetDir:     dst,
		tmpZipFile:    tmpZip,
	}

	err := inst.prepareZipFile(loading)
	if err != nil {
		return nil, err
	}

	err = inst.unzip(loading)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func (inst *zipTestingDirLoader) prepareZipFile(loading *zipTestingDirLoading) error {
	name := loading.sourceZipName
	data, err := loading.resources.GetBinary(name)
	if err != nil {
		return err
	}
	file := loading.tmpZipFile
	return file.GetIO().WriteBinary(data, nil, true)
}

func (inst *zipTestingDirLoader) unzip(loading *zipTestingDirLoading) error {
	src := loading.tmpZipFile
	dst := loading.targetDir
	reader, err := zip.OpenReader(src.Path())
	if err != nil {
		return err
	}
	defer reader.Close()
	items := reader.File
	for _, item := range items {

		target := dst.GetChild("./" + item.Name)
		err = inst.writeItemTo(target, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *zipTestingDirLoader) writeItemTo(to fs.Path, item *zip.File) error {
	info := item.FileInfo()
	if info.IsDir() {
		to.Mkdirs()
		return nil
	}
	// open reader
	reader, err := item.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	// open writer
	opt := fs.Options{Create: true}
	writer, err := to.GetIO().OpenWriter(&opt, true)
	if err != nil {
		return err
	}
	defer writer.Close()
	// pump
	buffer := make([]byte, 1024)
	for {
		cnt, _ := reader.Read(buffer)
		if cnt > 0 {
			writer.Write(buffer[0:cnt])
		} else {
			break
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type sparseTestingDirLoader struct {
}

func (inst *sparseTestingDirLoader) Load(rt TestingRuntime, name string, dst fs.Path) (fs.Path, error) {
	const r = true
	res := rt.Context().GetResources()
	items := res.List(name, r)
	for _, item := range items {
		target := dst.GetChild(item.RelativePath)
		err := inst.writeItem(item, res, target)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func (inst *sparseTestingDirLoader) writeItem(item *collection.Resource, res collection.Resources, target fs.Path) error {

	if item.IsDir {
		target.Mkdirs()
		return nil
	}

	data, err := res.GetBinary(item.AbsolutePath)
	if err != nil {
		return err
	}

	vlog.Debug("write testing data to file ", target.Path())
	return target.GetIO().WriteBinary(data, nil, true)
}

////////////////////////////////////////////////////////////////////////////////
