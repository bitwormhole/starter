package collection

import (
	"strings"
	"testing"

	"github.com/bitwormhole/starter/vlog"
)

func TestArgs(t *testing.T) {

	args := []string{"cmd1", "cmd2", "-a", "-b", "2", "-c", "31", "32", "--xxx", "777", "771", "772", "773", "---yyy=8", "888", "---zzz", "999"}
	arguments := CreateArguments()
	arguments.Import(args)
	reader := arguments.NewReader()

	keys := []string{"", "-a", "-b", "-c", "--xxx", "---yyy", "---zzz"}

	for _, key := range keys {
		flag := reader.GetFlag(key)
		if flag.Exists() {
			vlog.Info(flag.GetName())
			if strings.HasPrefix(flag.GetName(), "--") {
				f1, ok := flag.Pick(1)
				if ok {
					vlog.Info(" ..... padding: ", f1)
				}
			}
		} else {
			vlog.Info(" .................  want flag, but have not, name=", key)
		}
	}

	index := 0
	for {
		text, ok := reader.PickNext()
		if ok {
			vlog.Info("other.list[", index, "] = ", text)
		} else {
			break
		}
	}
	vlog.Info("done.")
}
