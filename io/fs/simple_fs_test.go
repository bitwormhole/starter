package fs

import (
	"testing"

	"github.com/bitwormhole/starter/vlog"
)

func TestGetParent(t *testing.T) {

	path := Default().GetPath(t.TempDir() + "/a/b/c/d")

	for ttl := 99; path != nil; {

		if ttl > 0 {
			ttl--
		} else {
			break
		}

		// pstr := path.Path()
		ext := path.Exists()
		isdir := path.IsDir()

		vlog.Debug("node path:[", path, "] exists:", ext, " isdir:", isdir)

		path = path.Parent()
	}
	vlog.Debug("node[done]")
}

func TestGetParent2(t *testing.T) {

	timeout := 100
	fs := Default()
	path := fs.GetPath("c:\\d\\e\\f\\g")

	if fs.SeparatorChar() == '/' {
		// for posix
		path = fs.GetPath("/a/b/c/x/y/z")
	}

	p := path
	for ; p != nil; p = p.Parent() {
		t.Log("path=", p.Path())
		if timeout < 0 {
			t.Error("timeout while call fs.Path.Parent()")
			break
		} else {
			timeout--
		}
	}
}
