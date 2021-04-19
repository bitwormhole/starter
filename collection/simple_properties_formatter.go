package collection

import (
	"sort"
	"strings"
)

type simplePropertiesFormatter struct {
	builder           strings.Builder
	enableSegment     bool
	currentSegmentKey string
}

// FormatProperties 将属性表格式化为文本形式
func FormatProperties(tab Properties) string {
	formatter := &simplePropertiesFormatter{}
	formatter.enableSegment = false
	formatter.handleProps(tab)
	return formatter.builder.String()
}

// FormatPropertiesWithSegment 将属性表格式化为文本形式(支持分段)
func FormatPropertiesWithSegment(tab Properties) string {
	formatter := &simplePropertiesFormatter{}
	formatter.enableSegment = true
	formatter.handleProps(tab)
	return formatter.builder.String()
}

func (inst *simplePropertiesFormatter) handleProps(props Properties) {

	table := props.Export(nil)
	keys := []string{}

	for key := range table {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for index := range keys {
		key := keys[index]
		val := table[key]
		inst.handleProperty(key, val)
	}
}

func (inst *simplePropertiesFormatter) handleProperty(name string, value string) {

	name = strings.TrimSpace(name)
	value = strings.TrimSpace(value)

	if len(name) == 0 || len(value) == 0 {
		return
	}

	if !inst.enableSegment {
		inst.builder.WriteString(name)
		inst.builder.WriteRune('=')
		inst.builder.WriteString(value)
		inst.builder.WriteRune('\n')
		return
	}

	inst.handlePropertyWithSegment(name, value)
}

func (inst *simplePropertiesFormatter) handlePropertyWithSegment(name string, value string) {

	n1, n2, n3 := inst.parsePropertyName(name)
	n1 = strings.TrimSpace(n1)
	n2 = strings.TrimSpace(n2)
	n3 = strings.TrimSpace(n3)
	segmentKey := n1 + "." + n2

	if segmentKey != inst.currentSegmentKey {
		// write segment header
		inst.currentSegmentKey = segmentKey
		if n2 == "" {
			inst.builder.WriteString("[")
			inst.builder.WriteString(n1)
			inst.builder.WriteString("]\n")
		} else {
			inst.builder.WriteString("[")
			inst.builder.WriteString(n1)
			inst.builder.WriteString(" \"")
			inst.builder.WriteString(n2)
			inst.builder.WriteString("\"]\n")
		}
	}

	// write k=v
	inst.builder.WriteString("  ")
	inst.builder.WriteString(n3)
	inst.builder.WriteRune('=')
	inst.builder.WriteString(value)
	inst.builder.WriteRune('\n')
}

func (inst *simplePropertiesFormatter) parsePropertyName(name string) (nameType string, nameID string, nameField string) {

	parts := strings.Split(name, ".")
	count := len(parts)

	if count <= 1 {
		// 'ab'
		return "", "", name

	} else if count == 2 {
		// 'a.b'
		return parts[0], "", parts[1]

	} else if count == 3 {
		// 'a.b.c'
		return parts[0], parts[1], parts[2]
	}

	// the nameID is formatted like 'm...n'
	var idBuilder strings.Builder
	iEnd := count - 1

	for i := 1; i < iEnd; i++ {
		if i > 1 {
			idBuilder.WriteRune('.')
		}
		idBuilder.WriteString(parts[i])
	}

	id := idBuilder.String()
	return parts[0], id, parts[iEnd]
}
