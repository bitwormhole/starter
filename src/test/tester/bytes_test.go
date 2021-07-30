package tester

import (
	"crypto/sha1"
	"fmt"
	"testing"

	"github.com/bitwormhole/starter/util"
)

func TestStringifyBytes(t *testing.T) {

	data := []byte{'h', 'e', 'l', 'l', 'o'}

	text := util.StringifyBytes(data)
	fmt.Println("data = ", text)

	sum := sha1.Sum(data)
	text = util.StringifyBytes(sum[:])
	fmt.Println("sha1sum(data) = ", text)

}
