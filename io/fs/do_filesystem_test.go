package fs

import (
	"strings"
	"testing"
)

func prepareDirForTest(t *testing.T) Path {
	fs := Default()
	dir := fs.GetPath(t.TempDir())

	dir.GetChild("chdir1").Mkdirs()
	dir.GetChild("chfile1").CreateFileWithSize(1024*16+666, nil)

	return dir
}

// tests for FS

func TestFsRoots(t *testing.T) {

	dir := prepareDirForTest(t)
	fs := dir.FileSystem()

	roots := fs.Roots()

	for idx := range roots {
		root := roots[idx]
		t.Log(root.Path())
	}

}

func TestFsGetPath(t *testing.T) {

	dir := prepareDirForTest(t)
	fs := dir.FileSystem()

	path1 := dir.GetChild("example/abc")
	path2 := fs.GetPath(path1.Path())

	p1 := path1.Path()
	p2 := path2.Path()

	if strings.Compare(p1, p2) == 0 {
		t.Log("ok, path=" + p1)
	} else {
		t.Error("path1 != path2")
	}
}

func TestFsSeparator(t *testing.T) {

	dir := prepareDirForTest(t)
	fs := dir.FileSystem()

	sep1 := fs.Separator()
	sep2 := fs.SeparatorChar()
	sep3 := fs.PathSeparatorChar()
	sep4 := fs.PathSeparator()

	t.Log("s1=" + sep1)
	t.Log("s2=" + string(sep2))
	t.Log("s3=" + string(sep3))
	t.Log("s4=" + sep4)
}
