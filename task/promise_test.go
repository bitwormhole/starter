package task

import (
	"testing"
	"time"

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

	done := false

	p.Then(func(result interface{}) {
		vlog.Info(result)
	}).Catch(func(err error) {
		vlog.Error(err)
	}).Finally(func() {
		vlog.Debug("done")
		done = true
	})

	for {
		time.Sleep(time.Second)
		if done {
			break
		}
	}
}
