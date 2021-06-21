package collection

import "strings"

func CreateArguments() Arguments {
	return &SimpleArguments{}
}

////////////////////////////////////////////////////////////////////////////////
// struct SimpleArguments

type SimpleArguments struct {
	args     []string
	mappings map[string]*simpleArgumentMapping
}

func (inst *SimpleArguments) getMappings() map[string]*simpleArgumentMapping {
	list := inst.args
	table := inst.mappings
	if list == nil {
		list = make([]string, 0)
		inst.args = list
	}
	if table == nil {
		table = make(map[string]*simpleArgumentMapping)
		for index := range list {
			item := list[index]
			key := strings.TrimSpace(item)
			mapping := &simpleArgumentMapping{}
			mapping.ref = list[index:]
			table[key] = mapping
		}
		inst.mappings = table
	}
	return table
}

func (inst *SimpleArguments) copyArray(src []string) []string {
	if src == nil {
		return []string{}
	}
	size := len(src)
	dst := make([]string, size)
	for index := range src {
		dst[index] = src[index]
	}
	return dst
}

func (inst *SimpleArguments) Export() []string {
	return inst.copyArray(inst.args)
}

func (inst *SimpleArguments) Import(args []string) {
	inst.args = inst.copyArray(args)
	inst.mappings = nil
}

func (inst *SimpleArguments) GetReader(flag string) (ArgumentReader, bool) {
	mappings := inst.getMappings()
	reader := &simpleArgumentReader{}
	mapping := mappings[flag]
	if mapping == nil {
		if flag == "" {
			mapping = &simpleArgumentMapping{}
			mapping.ref = inst.args
		} else {
			return reader, false
		}
	}
	reader.init(mapping)
	return reader, true
}

////////////////////////////////////////////////////////////////////////////////
// struct simpleArgumentMapping

type simpleArgumentMapping struct {
	ref []string
}

////////////////////////////////////////////////////////////////////////////////
// struct simpleArgumentReader

type simpleArgumentReader struct {
	ending string
	items  []string
	ptr    int
	size   int
	eof    bool
}

func (inst *simpleArgumentReader) init(mapping *simpleArgumentMapping) {

	if mapping == nil {
		return
	}
	items := mapping.ref
	if items == nil {
		return
	}
	inst.items = items
	inst.ptr = 0
	inst.ending = "-"
	inst.eof = false
	inst.size = len(items)
}

func (inst *simpleArgumentReader) SetEnding(ending string) {
	inst.ending = ending
}

func (inst *simpleArgumentReader) Ending() string {
	return inst.ending
}

func (inst *simpleArgumentReader) readNextItem() (string, bool) {
	if inst.eof {
		return "", false
	}
	items := inst.items
	ptr := inst.ptr
	size := inst.size
	for {
		if 0 <= ptr && ptr < size {
			text := items[ptr]
			ptr++
			if text == "" {
				continue
			}
			inst.ptr = ptr
			return text, true
		} else {
			inst.eof = true
			return "", false
		}
	}
}

func (inst *simpleArgumentReader) Read() (string, bool) {
	index := inst.ptr
	prefix := inst.ending
	item, ok := inst.readNextItem()
	if !ok {
		return "", false
	}
	if index > 0 {
		if strings.HasPrefix(item, prefix) {
			inst.eof = true
			return "", false
		}
	}
	return item, true
}
