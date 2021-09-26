package cli

import (
	"errors"
	"strings"
)

type taskListBuilder struct {
	list []*TaskUnit
}

func (inst *taskListBuilder) getList() []*TaskUnit {
	list := inst.list
	if list == nil {
		list = make([]*TaskUnit, 0)
		inst.list = list
	}
	return list
}

func (inst *taskListBuilder) create() []*TaskUnit {
	src := inst.getList()
	dst := make([]*TaskUnit, 0)
	for index, item := range src {
		item.LineNumber = index + 1
		args := item.Arguments
		if args == nil {
			continue
		}
		if len(args) == 0 {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

func (inst *taskListBuilder) addLine(line string, index int, args []string) error {

	unit := &TaskUnit{}
	unit.LineNumber = index + 1
	unit.Arguments = args
	unit.CommandLine = line

	list := inst.getList()
	list = append(list, unit)
	inst.list = list
	return nil
}

func (inst *taskListBuilder) parseLine(line string, index int) error {

	args := make([]string, 0)
	reader := &commandLineReader{}
	reader.init(line)

	for {
		reader.skipSpace()
		if reader.hasMore() {
			element, err := reader.read()
			if err != nil {
				return err
			}
			if element != "" {
				args = append(args, element)
			}
		} else {
			break
		}
	}

	return inst.addLine(line, index, args)
}

func (inst *taskListBuilder) parseScript(script string) error {
	const ch1 = "\r"
	const ch2 = "\n"
	script = strings.ReplaceAll(script, ch1, ch2)
	array := strings.Split(script, ch2)
	for index, line := range array {
		err := inst.parseLine(line, index)
		if err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type commandLineReader struct {
	source  []rune
	ptr     int
	length  int
	builder strings.Builder
}

func (inst *commandLineReader) init(line string) {
	array := ([]rune(line))
	inst.ptr = 0
	inst.length = len(array)
	inst.source = array
}

func (inst *commandLineReader) hasMore() bool {
	return inst.ptr < inst.length
}

func (inst *commandLineReader) read() (string, error) {

	i := inst.ptr
	length := inst.length
	array := inst.source

	if i >= length {
		return "", errors.New("EOF")
	}

	const mk1 = '\''
	const mk2 = '"'
	ch := array[i]

	if ch == '=' {
		inst.ptr = i + 1
		return "", nil
	} else if ch == mk1 {
		return inst.readText(mk1)
	} else if ch == mk2 {
		return inst.readText(mk2)
	} else {
		return inst.readToken()
	}
}

func (inst *commandLineReader) readToken() (string, error) {

	builder := &inst.builder
	i := inst.ptr
	length := inst.length
	array := inst.source
	chs := []rune{' ', '\t', '=', '"', '\'', '\n', '\r'} // ending of token

	builder.Reset()

	for ; i < length; i++ {
		ch := array[i]
		if inst.containsRune(ch, chs) {
			break
		}
		builder.WriteRune(ch)
	}

	inst.ptr = i
	return builder.String(), nil
}

func (inst *commandLineReader) readText(ending rune) (string, error) {

	builder := &inst.builder
	i := inst.ptr + 1
	length := inst.length
	array := inst.source

	builder.Reset()

	for ; i < length; i++ {
		ch := array[i]
		if ch == ending {
			inst.ptr = i + 1
			return builder.String(), nil
		}
		builder.WriteRune(ch)
	}

	return "", errors.New("end with exception")
}

func (inst *commandLineReader) skipSpace() {
	i := inst.ptr
	length := inst.length
	array := inst.source
	chs := []rune{' ', '\t', '\n', '\r'}
	for ; i < length; i++ {
		ch := array[i]
		if inst.containsRune(ch, chs) {
			continue
		} else {
			break
		}
	}
	inst.ptr = i
}

func (inst *commandLineReader) containsRune(ch rune, list []rune) bool {
	if list == nil {
		return false
	}
	for _, ch2 := range list {
		if ch == ch2 {
			return true
		}
	}
	return false
}
