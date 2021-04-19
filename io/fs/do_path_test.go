package fs

import (
	"fmt"
	"strings"
	"testing"
)

// tests for Path

func TestPathName(t *testing.T) {

	dir := prepareDirForTest(t)

	name1 := "test_path.name"
	path := dir.GetChild(name1)
	name2 := path.Name()

	if strings.Compare(name1, name2) == 0 {
		t.Log("ok, name=" + name2)
	} else {
		t.Error("name1 != name2")
	}
}

func TestPath2(t *testing.T) {

	dir := prepareDirForTest(t)

	path1 := dir.GetChild("test//.//path2").Path()
	path2 := dir.GetChild("test\\path2").Path()

	if strings.Compare(path1, path2) == 0 {
		t.Log("ok, path=" + path1)
	} else {
		t.Error("path1 != path2")
	}
}

func TestPathExists(t *testing.T) {

	dir := prepareDirForTest(t)
	file := dir.GetChild("test.path.exists")

	if dir.Exists() && !file.Exists() {
		t.Log("ok")
		t.Log("exists:      dir = " + dir.Path())
		t.Log("not exists: file = " + file.Path())
	} else {
		t.Error("bad path.exists")
	}
}

func TestPathIsDir(t *testing.T) {

	dir := prepareDirForTest(t)
	file := dir.GetChild("test.path.isdir")

	file.CreateFile(nil)

	if !dir.IsDir() {
		t.Error("error: dir is not dir")
	}

	if file.IsDir() {
		t.Error("error: file is dir")
	}
}

func TestPathIsFile(t *testing.T) {

	dir := prepareDirForTest(t)
	file := dir.GetChild("test.path.isfile")

	file.CreateFile(nil)

	if dir.IsFile() {
		t.Error("error: dir is file")
	}

	if !file.IsFile() {
		t.Error("error: file is not file")
	}
}

func TestPathFS(t *testing.T) {
	dir := prepareDirForTest(t)
	fs := dir.FileSystem()
	if fs == nil {
		t.Error("error: fs == nil ")
	}
}

func TestPathParent(t *testing.T) {

	dir0 := prepareDirForTest(t)
	dir1 := dir0.GetChild("a")
	dir2 := dir1.GetChild("b")
	dir3 := dir2.GetChild("c")

	p3 := dir3.Parent().Parent().Parent().Path()
	p2 := dir2.Parent().Parent().Path()
	p1 := dir1.Parent().Path()
	p0 := dir0.Path()

	list := []string{p0, p1, p2, p3}

	for i := range list {
		str := list[i]
		if strings.Compare(p0, str) == 0 {
			continue
		} else {
			t.Error("bad path: " + str)
		}
	}

	ptr := dir0
	findRoot := false

	for limit := 30; limit > 0; limit-- {
		if ptr == nil {
			findRoot = true
			break
		} else {
			ptr = ptr.Parent()
		}
	}

	if !findRoot {
		t.Error("no root at path: " + dir0.Path())
	}
}

func TestPathSize(t *testing.T) {
	size1 := int64(1234)
	dir := prepareDirForTest(t)
	file := dir.GetChild("test.pathSize")
	err := file.CreateFileWithSize(size1, nil)

	if err != nil {
		t.Error("cannot create file: " + file.Path())
	}

	size2 := file.Size()

	if size1 == size2 {
		t.Log("ok, file.size = ", size2)
	} else {
		t.Error("bad file size: ", size2)
	}
}

func TestPathLastModTime(t *testing.T) {
	dir := prepareDirForTest(t)
	time := dir.LastModTime()
	t.Log("dir.lastModTime = " + fmt.Sprint(time))
}

func TestPathCreateFile(t *testing.T) {

	dir := prepareDirForTest(t)
	file := dir.GetChild("test.path.createfile")
	err := file.CreateFile(nil)

	if err != nil {
		t.Error("cannot create file: ", err)
	}

	if !file.IsFile() {
		t.Error("file not exists")
	}

	if file.Size() != 0 {
		t.Error("bad file size")
	}
}

func TestPathCreateFileWithSize(t *testing.T) {

	size := 12345
	dir := prepareDirForTest(t)
	file := dir.GetChild("test.path.createfile")
	err := file.CreateFileWithSize(int64(size), nil)

	if err != nil {
		t.Error("cannot create file: ", err)
	}

	if !file.IsFile() {
		t.Error("file not exists")
	}

	if file.Size() != int64(size) {
		t.Error("bad file size")
	}
}

func TestPathMkdir(t *testing.T) {

	dir := prepareDirForTest(t)
	dir2 := dir.GetChild("test.path.mkdir")

	err := dir2.Mkdir()
	if err != nil {
		t.Error(err)
		return
	}

	if !dir2.IsDir() {
		t.Error("error: no mkdir, dir=" + dir2.Path())
	}
}

func TestPathMkdirs(t *testing.T) {

	dir := prepareDirForTest(t)
	dir2 := dir.GetChild("test/.path/.mkdir")

	err := dir2.Mkdirs()
	if err != nil {
		t.Error(err)
		return
	}

	if !dir2.IsDir() {
		t.Error("error: no mkdirs, dir=" + dir2.Path())
	}
}

func TestPathGetMoveTo(t *testing.T) {

	dir := prepareDirForTest(t)
	file1 := dir.GetChild("test.path.moveto.file1")
	file2 := dir.GetChild("test.path.moveto.file2")

	size1 := int64(123)
	file1.CreateFileWithSize(size1, nil)
	file1.MoveTo(file2)
	size2 := file2.Size()

	if size2 != size1 {
		t.Error("bad file2.size: ", size2)
	}

}

func TestPathCopyTo(t *testing.T) {

	dir := prepareDirForTest(t)
	file1 := dir.GetChild("test.path.copyto.file1")
	file2 := dir.GetChild("test.path.copyto.file2")

	size1 := int64(1234)
	file1.CreateFileWithSize(size1, nil)
	file1.CopyTo(file2)
	size2 := file2.Size()

	if size2 != size1 {
		t.Error("bad file2.size: ", size2)
	}

}

func TestPathDelete(t *testing.T) {

	dir := prepareDirForTest(t)
	dir2 := dir.GetChild("test.path/.delete")

	err := dir2.Mkdirs()
	if err != nil {
		t.Error(err)
	}

	if !dir2.IsDir() {
		t.Error("bad dir status")
	}

	err = dir2.Delete()
	if err != nil {
		t.Error(err)
	}

	if dir2.Exists() {
		t.Error("bad dir status")
	}

}

func TestPathGetNameList(t *testing.T) {

	dir := prepareDirForTest(t)
	list := dir.GetNameList()

	for index := range list {
		item := list[index]
		t.Log("GetNameList.item: ", item)
	}

}

func TestPathGetPathList(t *testing.T) {

	dir := prepareDirForTest(t)
	list := dir.GetPathList()

	for index := range list {
		item := list[index]
		t.Log("GetNameList.item: ", item)
	}

}

func TestPathGetItemList(t *testing.T) {

	dir := prepareDirForTest(t)
	list := dir.GetItemList()

	for index := range list {
		item := list[index]
		t.Log("GetNameList.item: ", item)
	}

}

func TestPathGetHref(t *testing.T) {

	dir := prepareDirForTest(t)
	file := dir.GetChild("test.path.gethref")
	file.CreateFile(nil)

	xyz1 := dir.GetChild("xyz")
	xyz2 := file.GetHref("./xyz")

	p1 := xyz1.Path()
	p2 := xyz2.Path()

	if strings.Compare(p1, p2) == 0 {
		t.Log("ok, path = " + xyz1.Path())
	} else {
		t.Error("path1 != path2")
	}
}

func TestPathGetChild(t *testing.T) {
	dir := prepareDirForTest(t)
	file := dir.GetChild("./a/b/c")
	c := file.Name()
	b := file.Parent().Name()

	if c != "c" && b != "b" {
		t.Error("bad child path (GetChild): " + file.Path())
	}
}
