package collection

import (
	"errors"
)

// SimpleProperties 实现一个简单的 Properties 集合
type SimpleProperties struct {
	table map[string]string
}

func (inst *SimpleProperties) getTable() map[string]string {
	t := inst.table
	if t == nil {
		t = make(map[string]string)
		inst.table = t
	}
	return t
}

// GetPropertyRequired 在集合中取指定的值，如果没有则返回error
func (inst *SimpleProperties) GetPropertyRequired(name string) (string, error) {
	t := inst.getTable()
	value := t[name]
	if value == "" {
		return "", errors.New("no property with name: " + name)
	}
	return value, nil
}

// GetProperty 在集合中取指定的值，如果没有则返回默认值
func (inst *SimpleProperties) GetProperty(name string, defaultValue string) string {
	t := inst.getTable()
	value := t[name]
	if value == "" {
		return defaultValue
	}
	return value
}

// SetProperty 给集合中指定的键值对赋值
func (inst *SimpleProperties) SetProperty(name string, value string) {
	t := inst.getTable()
	t[name] = value
}

// Clear 清空本集合中的所有内容
func (inst *SimpleProperties) Clear() {
	inst.table = make(map[string]string)
}

// Export 把本集合中的所有键值对导出
func (inst *SimpleProperties) Export(dst map[string]string) map[string]string {

	src := inst.getTable()

	if dst == nil {
		dst = make(map[string]string)
	}

	for key := range src {
		dst[key] = src[key]
	}

	return dst
}

// Import 把参数 src 中的所有键值对导入本集合，保留原有内容（如果没有被覆盖）
func (inst *SimpleProperties) Import(src map[string]string) {

	if src == nil {
		return
	}

	dst := inst.getTable()

	for key := range src {
		dst[key] = src[key]
	}
}

func CreateProperties() Properties {
	return &SimpleProperties{}
}
