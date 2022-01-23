package util

import "encoding/base64"

// Base64 表示字符串形式的 base64 值
type Base64 string

// String 转换为base64格式的字符串
func (b Base64) String() string {
	return string(b)
}

// Bytes 转换为 bytes
func (b Base64) Bytes() []byte {
	str := b.String()
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return []byte{}
	}
	return data
}

// HexString 转换为16进制格式的字符串
func (b Base64) HexString() Hex {
	data := b.Bytes()
	return HexFromBytes(data)
}

////////////////////////////////////////////////////////////////////////////////

// Base64FromString 转换 b64 string ==> base64
func Base64FromString(s string) Base64 {
	return Base64(s)
}

// Base64FromBytes 转换 []byte ==> base64
func Base64FromBytes(b []byte) Base64 {
	str := base64.StdEncoding.EncodeToString(b)
	return Base64(str)
}

// Base64FromHexString 转换 hex string ==> base64
func Base64FromHexString(hex Hex) Base64 {
	data := hex.Bytes()
	b64 := Base64FromBytes(data)
	return b64
}
