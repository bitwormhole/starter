package util

import (
	"crypto/sha1"
	"testing"
)

func TestHex(t *testing.T) {

	sum := sha1.Sum([]byte("12345678"))

	hex1 := HexFromBytes(sum[:])
	bytes1 := hex1.Bytes()
	str1 := hex1.String()

	hex2 := HexFromBytes(bytes1)
	hex3, err := HexFromString(str1)

	if err != nil {
		t.Error(err)
	}

	if hex1 != hex2 || hex2 != hex3 {
		t.Error("hex1 != hex2 != hex3")
	}
}
