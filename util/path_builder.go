package util

import (
	"errors"
	"strings"
)

type PathBuilder struct {
	inner *innerPathBuilder
}

type innerPathBuilder struct {
	parts           []string
	separator       string
	enableDoubleDot bool
	enableRoot      bool
	enableTrim      bool
}

//// inner

func (inst *innerPathBuilder) init() {
	inst.parts = make([]string, 0)
	inst.separator = "/"
	inst.enableDoubleDot = false
	inst.enableRoot = false
	inst.enableTrim = false
}

//// private

func (inst *PathBuilder) _inner() *innerPathBuilder {
	inner := inst.inner
	if inner == nil {
		inner = &innerPathBuilder{}
		inner.init()
		inst.inner = inner
	}
	return inner
}

func (inst *PathBuilder) _append(str string) {
	inner := inst._inner()
	inner.parts = append(inner.parts, str)
}

func (inst *PathBuilder) _pop_end(list []string) ([]string, error) {
	size := len(list)
	if size < 1 {
		return nil, errors.New("len(list)==0")
	}
	return list[0 : size-1], nil
}

//// public

func (inst *PathBuilder) SetSeparator(sep string) *PathBuilder {
	inner := inst._inner()
	inner.separator = sep
	return inst
}

func (inst *PathBuilder) EnableRoot(enable bool) *PathBuilder {
	inner := inst._inner()
	inner.enableRoot = enable
	return inst
}

func (inst *PathBuilder) EnableTrim(enable bool) *PathBuilder {
	inner := inst._inner()
	inner.enableTrim = enable
	return inst
}

func (inst *PathBuilder) EnableDoubleDot(enable bool) *PathBuilder {
	inner := inst._inner()
	inner.enableDoubleDot = enable
	return inst
}

func (inst *PathBuilder) AppendSimpleName(simpleName string) *PathBuilder {
	inst._append(simpleName)
	return inst
}

func (inst *PathBuilder) AppendSimpleNameList(list []string) *PathBuilder {
	if list == nil {
		return inst
	}
	for index := range list {
		item := list[index]
		inst._append(item)
	}
	return inst
}

func (inst *PathBuilder) AppendPath(path string) *PathBuilder {
	inner := inst._inner()
	if inner.enableRoot {
		if strings.HasPrefix(path, "/") {
			inner.parts = make([]string, 0)
		}
	}
	path = strings.ReplaceAll(path, "\\", "/")
	array := strings.Split(path, "/")
	return inst.AppendSimpleNameList(array)
}

func (inst *PathBuilder) String() string {
	str, _ := inst.Create("", "")
	return str
}

func (inst *PathBuilder) Create(prefix string, suffix string) (string, error) {

	inner := inst._inner()
	list1 := inner.parts
	list2 := make([]string, 0)
	en_trim := inner.enableTrim
	en_dotdot := inner.enableDoubleDot

	// filter items
	for index := range list1 {
		item := list1[index]
		if en_trim {
			item = strings.TrimSpace(item)
		}
		if item == "" {
			continue // skip
		} else if item == "." {
			continue // skip
		} else if item == ".." {
			if en_dotdot {
				list22, err := inst._pop_end(list2)
				if err != nil {
					return "", err
				}
				list2 = list22
			} else {
				return "", errors.New("disable name:[..]")
			}
		} else {
			list2 = append(list2, item)
		}
	}

	// to string
	sep := ""
	builder := &strings.Builder{}
	builder.WriteString(prefix)
	for _, item := range list2 {
		builder.WriteString(sep)
		builder.WriteString(item)
		sep = inner.separator
	}
	builder.WriteString(suffix)
	return builder.String(), nil
}
