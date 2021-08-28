package util

import (
	"encoding/json"
	"errors"
	"strings"
)

// ErrorBuilder 用来创建一个错误对象
type ErrorBuilder struct {
	message    string
	inner      error
	properties map[string]string
}

// Message 设置错误消息
func (inst *ErrorBuilder) Message(msg string) *ErrorBuilder {
	inst.message = msg
	return inst
}

// InnerError 设置内嵌错误
func (inst *ErrorBuilder) InnerError(inner error) *ErrorBuilder {
	inst.inner = inner
	return inst
}

// Set 设置属性
func (inst *ErrorBuilder) Set(name string, value string) *ErrorBuilder {
	table := inst.properties
	if table == nil {
		table = make(map[string]string)
		inst.properties = table
	}
	table[name] = value
	return inst
}

// Create 创建错误对象
func (inst *ErrorBuilder) Create() error {
	table := inst.properties
	inner := inst.inner
	builder := strings.Builder{}
	builder.WriteString(inst.message)
	if table != nil {
		data, err := json.Marshal(table)
		if err == nil {
			builder.WriteString(":")
			builder.WriteString(string(data))
		} else {
			builder.WriteString(", error:")
			builder.WriteString(err.Error())
		}
	}
	if inner != nil {
		builder.WriteString(", inner:")
		builder.WriteString(inner.Error())
	}
	return errors.New(builder.String())
}
