package lang

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringifyObject(t *testing.T) {

	doTestStringifyObject(nil)

	doTestStringifyObject(t)
	doTestStringifyObject(false)
	doTestStringifyObject(true)
	doTestStringifyObject(0)
	doTestStringifyObject(1)
	doTestStringifyObject(3.14)
	doTestStringifyObject("666")

	doTestStringifyObject(&strings.Builder{})
	doTestStringifyObject(CreateReleasePool())
	doTestStringifyObject(CreateReleasePool)
}

func doTestStringifyObject(o Object) {
	str := StringifyObject(o)
	fmt.Println("StringifyObject: ", str)
}
