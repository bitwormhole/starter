package util

import (
	"bytes"
	"errors"
	"strings"
)

// private
func stringifyBytesDigit(n int, builder *strings.Builder) {
	n = n & 0x0f
	if (0 <= n) && (n <= 9) {
		builder.WriteRune('0' + rune(n))
	} else {
		builder.WriteRune('a' + rune(n-10))
	}
}

// StringifyBytes 是 ToHexString 的别名
func StringifyBytes(data []byte) string {
	return ToHexString(data)
}

// ToHexString 把字节数组格式化为字符串（charset=0~9+abcdef）
func ToHexString(data []byte) string {
	if data == nil {
		return ""
	}
	builder := strings.Builder{}
	for index := range data {
		b := int(data[index])
		stringifyBytesDigit(b>>4, &builder)
		stringifyBytesDigit(b, &builder)
	}
	return builder.String()
}

// ParseHexString 解析hex形式的字符串，返回对应的 []byte
func ParseHexString(hex string) ([]byte, error) {
	builder := bytes.Buffer{}
	chs := []rune(hex)
	var err error
	countHex := 0
	n0 := 0
	for _, ch := range chs {
		n := 0
		if '0' <= ch && ch <= '9' {
			n = int(ch - '0')
		} else if 'a' <= ch && ch <= 'f' {
			n = int(ch-'a') + 0x0a
		} else if 'A' <= ch && ch <= 'F' {
			n = int(ch-'A') + 0x0a
		} else {
			if err == nil {
				err = errors.New("bad hex string: [" + hex + "]")
			}
			continue
		}
		if (countHex & 0x01) == 0 {
			n0 = n << 4
		} else {
			b := (n0 & 0xf0) | (n & 0x0f)
			builder.WriteByte(byte(b))
		}
		countHex++
	}
	data := builder.Bytes()
	return data, err
}
