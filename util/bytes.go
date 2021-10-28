package util

import "strings"

func stringifyBytesDigit(n int, builder *strings.Builder) {
	n = n & 0x0f
	if (0 <= n) && (n <= 9) {
		builder.WriteRune('0' + rune(n))
	} else {
		builder.WriteRune('a' + rune(n-10))
	}
}

// StringifyBytes 把字节数组格式化为字符串（charset=0~9+abcdef）
func StringifyBytes(data []byte) string {
	if data == nil {
		return "[nil]"
	}
	builder := strings.Builder{}
	for index := range data {
		b := int(data[index])
		stringifyBytesDigit(b>>4, &builder)
		stringifyBytesDigit(b, &builder)
	}
	return builder.String()
}
