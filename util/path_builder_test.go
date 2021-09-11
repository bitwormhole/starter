package util

import (
	"fmt"
	"testing"
)

func TestPathBuilder(t *testing.T) {

	builder := &PathBuilder{}
	builder.EnableRoot(true)
	builder.EnableTrim(true)
	builder.IgnoreDoubleDot(false)

	builder.AppendPath("abc/def/g")
	builder.AppendPath("/   hijk ////")
	builder.AppendPath("lm/n")
	builder.AppendPath("o/..//p///q")

	path, _ := builder.Create("[", "]")
	fmt.Println("path=", path)
}
