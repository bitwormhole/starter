package lang

// Object 对象：相当于 interface{}
type Object interface {
}

// BaseObject 基本对象：比Object更复杂一点点
type BaseObject interface {
	Stringer
	Equals(other BaseObject) bool
	HashCode() int
}
