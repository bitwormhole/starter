package buffer

import (
	"testing"

	"github.com/bitwormhole/starter/io/fs"
)

func TestTextFileBuffer(t *testing.T) {

	dir := fs.Default().GetPath(t.TempDir())
	file := dir.GetChild("TestTextFileBuffer.txt")
	buffer := &TextFileBuffer{}
	buffer.Init(file)

	a1 := []string{"", "aaa", "bbb", "ccc", "ddd", "eee"}
	a2 := make([]string, 0)

	for _, item := range a1 {
		a2 = append(a2, item)
		a2 = append(a2, item)
		a2 = append(a2, item)
	}

	for _, item := range a2 {

		older := buffer.GetText(false)

		next := item
		buffer.SetText(next, false)
		n1 := buffer.GetText(false)
		n2 := buffer.GetText(false)

		t.Log("older = ", older)
		t.Log("=======================")
		t.Log(" next = ", next)
		t.Log("   n1 = ", n1)
		t.Log("   n2 = ", n2)

		if (next != n1) || (next != n2) {
			t.Fatal("(next != n1) || (next != n2)")
		}
	}
}
