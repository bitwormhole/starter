package collection

import "errors"

type SimpleEnvironment struct {
	table map[string]string
}

func (inst *SimpleEnvironment) GetEnv(name string) (string, error) {
	value := inst.table[name]
	if value == "" {
		return "", errors.New("no env: " + name)
	}
	return value, nil
}

func (inst *SimpleEnvironment) SetEnv(name string, value string) {
	inst.table[name] = value
}

func (inst *SimpleEnvironment) Export(dst map[string]string) map[string]string {
	if dst == nil {
		dst = make(map[string]string)
	}
	src := inst.table
	for key := range src {
		dst[key] = src[key]
	}
	return dst
}

func (inst *SimpleEnvironment) Import(src map[string]string) {
	if src == nil {
		return
	}
	dst := inst.table
	for key := range src {
		dst[key] = src[key]
	}
}

func CreateEnvironment() Environment {
	t := make(map[string]string)
	return &SimpleEnvironment{table: t}
}
