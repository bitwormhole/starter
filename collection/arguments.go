package collection

// Arguments 代表参数表对象
type Arguments interface {
	NewReader() ArgumentReader
	Length() int
	Get(index int) string
	Import([]string)
	Export() []string
}

// ArgumentReader 代表一个参数表的读者
type ArgumentReader interface {
	Flags() []string
	GetFlag(name string) ArgumentFlag
	PickNext() (string, bool)
}

// ArgumentFlag 代表参数表中的一个标志
type ArgumentFlag interface {
	GetName() string
	Exists() bool
	Pick(offset int) (string, bool)
}
