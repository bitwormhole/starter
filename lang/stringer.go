package lang

// Stringer 是一个简单的接口，它把对象格式化为字符串
type Stringer interface {
	String() string
}
