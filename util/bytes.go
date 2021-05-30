package util

import "strings"

func _stringifyBytes_digit(n int, builder *strings.Builder) {
	n = n & 0x0f
	if (0 <= n) && (n <= 9) {
		builder.WriteRune('0' + rune(n))
	} else {
		builder.WriteRune('a' + rune(n-10))
	}
}

func StringifyBytes(data []byte) string {
	if data == nil {
		return "[nil]"
	}
	builder := &strings.Builder{}
	for index := range data {
		b := int(data[index])
		_stringifyBytes_digit(b>>4, builder)
		_stringifyBytes_digit(b, builder)
	}
	return builder.String()
}
