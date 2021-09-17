package task

import (
	"testing"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

func TestPromise(t *testing.T) {

	dir := t.TempDir()
	file := fs.Default().GetPath(dir).GetChild("test.p")

	// file.GetIO().WriteText("hello,world", nil, true)

	p := NewPromise(func(resolve ResolveFn, reject RejectFn) {
		txt, err := file.GetIO().ReadText(nil)
		if err != nil {
			reject(err)
		} else {
			resolve(txt)
		}
	})

	p.Then(func(result interface{}) {
		vlog.Info(result)
	}).Catch(func(err error) {
		vlog.Error(err)
	}).Finally(func() {
		vlog.Debug("done")
	})
}
