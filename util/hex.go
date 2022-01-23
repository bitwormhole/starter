package util

import (
	"errors"
	"strings"
)

// Hex 表示十六进制形式的字符串
type Hex string

// Bytes 这个函数把 Hex 的值转换成字节数组
func (h Hex) Bytes() []byte {
	str := string(h)
	data, err := ParseHexString(str)
	if err != nil {
		return []byte{}
	}
	return data
}

// String 这个函数把 Hex 的值转换成字符串
func (h Hex) String() string {
	return string(h)
}

////////////////////////////////////////////////////////////////////////////////

// HexFromString 把字符串转换成 hex
func HexFromString(s string) (Hex, error) {
	builder := strings.Builder{}
	chs := []rune(s)
	var err error
	for _, ch := range chs {
		if '0' <= ch && ch <= '9' {
			builder.WriteRune(ch)
		} else if 'a' <= ch && ch <= 'f' {
			builder.WriteRune(ch)
		} else if 'A' <= ch && ch <= 'F' {
			n := ch - 'A'
			ch = 'a' + n
			builder.WriteRune(ch)
		} else {
			if err == nil {
				err = errors.New("bad hex string: " + s)
			}
		}
	}
	str := builder.String()
	return Hex(str), err
}

// HexFromBytes 把字节数组转换成 hex
func HexFromBytes(b []byte) Hex {
	str := StringifyBytes(b)
	return Hex(str)
}
