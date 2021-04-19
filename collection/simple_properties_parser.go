package collection

import (
	"errors"
	"strings"
)

type simplePropertiesParser struct {
	output     Properties
	err        error
	errLineNum int
	keyPrefix  string
}

// ParseProperties 函数把参数 text 解析为属性表，存入dest中。
func ParseProperties(text string, container Properties) (Properties, error) {

	if container == nil {
		container = &SimpleProperties{}
	}

	parser := &simplePropertiesParser{}
	parser.output = container
	parser.parse(text)
	return container, parser.err
}

func (inst *simplePropertiesParser) parse(text string) {

	text = strings.ReplaceAll(text, "\r", "\n")
	lines := strings.Split(text, "\n")

	for index := range lines {
		line := strings.TrimSpace(lines[index])
		err := inst.handleLine(line)
		if err != nil {
			inst.err = err
			inst.errLineNum = index + 1
			break
		}
	}

}

func (inst *simplePropertiesParser) handleLine(line string) error {

	if len(line) == 0 {
		return nil
	}

	if strings.HasPrefix(line, "#") {
		return nil
	}

	if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
		return inst.handleSegment(line)
	} else {
		return inst.handleKeyValue(line)
	}
}

func (inst *simplePropertiesParser) handleSegment(line string) error {

	const quotation string = "\""
	line = strings.ReplaceAll(line, "'", quotation)

	i1 := strings.IndexRune(line, '[')
	i2 := strings.Index(line, quotation)
	i3 := strings.LastIndex(line, quotation)
	i4 := strings.IndexRune(line, ']')

	line = strings.ReplaceAll(line, "[", quotation)
	line = strings.ReplaceAll(line, "]", quotation)
	parts := strings.Split(line, quotation)
	part1 := ""
	part2 := ""

	if (0 <= i1) && (i1 < i2) && (i2 < i3) && (i3 < i4) {
		part1 = parts[1]
		part2 = parts[2]
	} else if (0 <= i1) && (i1 < i4) && (i2 < 0) && (i3 < 0) {
		part1 = parts[1]
	} else {
		return errors.New("bad segment line: " + line)
	}

	part1 = strings.TrimSpace(part1)
	part2 = strings.TrimSpace(part2)

	if len(part1) > 0 && len(part2) > 0 {
		inst.keyPrefix = part1 + "." + part2 + "."
	} else if len(part1) > 0 && len(part2) == 0 {
		inst.keyPrefix = part1 + "."
	} else {
		inst.keyPrefix = ""
	}

	return nil
}

func (inst *simplePropertiesParser) handleKeyValue(line string) error {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return errors.New("the line is not in format of 'key=value'")
	}
	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])
	if len(key) == 0 || len(val) == 0 {
		return nil
	}
	inst.output.SetProperty(inst.keyPrefix+key, val)
	return nil
}
