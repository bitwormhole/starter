package collection

import "strings"

// CreateArguments 创建一个空的参数表
func CreateArguments() Arguments {
	pNew := &innerSimpleArguments{}
	pNew.init(nil)
	return pNew
}

// InitArguments 用给定的数据初始化一个新的参数表
func InitArguments(args []string) Arguments {
	pNew := &innerSimpleArguments{}
	pNew.init(args)
	return pNew
}

////////////////////////////////////////////////////////////////////////////////

// struct innerSimpleArguments
type innerSimpleArguments struct {
	args []string
	size int
}

func (inst *innerSimpleArguments) init(args []string) Arguments {
	inst.args = inst.makeCopy(args)
	inst.size = len(args)
	return inst
}

func (inst *innerSimpleArguments) Length() int {
	return inst.size
}

func (inst *innerSimpleArguments) Get(index int) string {
	return inst.args[index]
}

func (inst *innerSimpleArguments) NewReader() ArgumentReader {
	args := inst.makeCopy(inst.args)
	reader := &innerArgumentReader{}
	return reader.init(args)
}

func (inst *innerSimpleArguments) Import(args []string) {
	inst.args = inst.makeCopy(args)
}

func (inst *innerSimpleArguments) Export() []string {
	return inst.makeCopy(inst.args)
}

func (inst *innerSimpleArguments) makeCopy(src []string) []string {
	if src == nil {
		src = []string{}
	}
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

////////////////////////////////////////////////////////////////////////////////

// struct innerSimpleArguments
type innerArgumentReader struct {
	args    []string
	flags   map[string]ArgumentFlag
	flagNON ArgumentFlag
	ptr     int
	size    int
}

func (inst *innerArgumentReader) init(args []string) ArgumentReader {
	flags := make(map[string]ArgumentFlag)
	for index, element := range args {
		if strings.HasPrefix(element, "-") {
			flag := &innerArgumentFlag{}
			flag.index = index
			flag.name = element
			flag.reader = inst
			flags[element] = flag
			args[index] = ""
		}
	}
	inst.flagNON = &innerArgumentFlag{}
	inst.flags = flags
	inst.args = args
	inst.ptr = 0
	inst.size = len(args)
	return inst
}

func (inst *innerArgumentReader) Flags() []string {
	src := inst.flags
	dst := make([]string, 0)
	for name := range src {
		dst = append(dst, name)
	}
	return dst
}

func (inst *innerArgumentReader) GetFlag(name string) ArgumentFlag {
	flag := inst.flags[name]
	if flag == nil {
		flag = inst.flagNON
	}
	return flag
}

func (inst *innerArgumentReader) PickNext() (string, bool) {
	args := inst.args
	i := inst.ptr
	size := inst.size
	text := ""
	ok := false
	for ; i < size; i++ {
		el := args[i]
		if el != "" {
			text = el
			ok = true
			i++
			break
		}
	}
	inst.ptr = i
	return text, ok
}

////////////////////////////////////////////////////////////////////////////////

// struct innerSimpleArguments
type innerArgumentFlag struct {
	reader *innerArgumentReader
	name   string
	index  int
}

func (inst *innerArgumentFlag) _Impl() ArgumentFlag {
	return inst
}

func (inst *innerArgumentFlag) GetName() string {
	return inst.name
}

func (inst *innerArgumentFlag) Exists() bool {
	return inst.reader != nil
}

func (inst *innerArgumentFlag) Pick(offset int) (string, bool) {
	reader := inst.reader
	if reader == nil {
		return "", false
	}
	args := reader.args
	index := inst.index + offset
	if 0 <= index && index < reader.size {
		text := args[index]
		if text != "" {
			args[index] = ""
			return text, true
		}
	}
	return "", false
}

////////////////////////////////////////////////////////////////////////////////
