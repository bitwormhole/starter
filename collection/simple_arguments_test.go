package collection

import (
	"fmt"
	"testing"
)

func TestArgs(t *testing.T) {

	args := []string{"cmd1", "cmd2", "-a", "-b", "2", "-c", "31", "32", "--xxx", "777", "---yyy=8", "888", "---zzz", "999"}
	arguments := CreateArguments()
	arguments.Import(args)

	keys := []string{"", "-a", "-b", "-c", "--xxx", "---yyy", "---zzz"}

	for index := range keys {
		key := keys[index]
		fmt.Println("Read:", key)
		reader, ok := arguments.GetReader(key)
		if key == "" {
			reader.SetEnding("!")
		}
		if ok {
			for {
				item, ok2 := reader.Read()
				if ok2 {
					fmt.Println("  \\-item:", item)
				} else {
					break
				}
			}
		}
	}
}
