package util

import (
	"crypto/sha1"
	"testing"
)

func TestBase64(t *testing.T) {

	sum := sha1.Sum([]byte("12345678"))

	base1 := Base64FromBytes(sum[:])
	str1 := base1.String()
	hex1 := base1.HexString()
	bytes1 := base1.Bytes()

	base2 := Base64FromHexString(hex1)
	base3 := Base64FromString(str1)
	base4 := Base64FromBytes(bytes1)

	if base1 != base2 || base2 != base3 || base3 != base4 {
		t.Error("base1 != base2 != base3 != base4")
	}
}
