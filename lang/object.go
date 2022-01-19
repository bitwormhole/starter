package lang

import (
	"reflect"
)

// Object 对象：相当于 interface{}
type Object interface {
}

// BaseObject 基本对象：比Object更复杂一点点
type BaseObject interface {
	Stringer
	Equals(other BaseObject) bool
	HashCode() int
}

// StringifyObject 生成对象的摘要字符串， 类似于 java.lang.Object.toString()
func StringifyObject(o interface{}) string {
	if o == nil {
		return "[nil]"
	}
	t := reflect.TypeOf(o)

	// v := reflect.ValueOf(o)
	// return fmt.Sprint(t.String(), "(", v.Pointer(), ")")

	return t.String() + "{}"
}
