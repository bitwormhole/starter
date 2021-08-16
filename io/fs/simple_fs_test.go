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
}
