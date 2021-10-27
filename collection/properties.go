package collection

// Properties 接口表示对属性列表的引用。
type Properties interface {
	GetPropertyRequired(name string) (string, error)
	GetProperty(name string, defaultValue string) string
	SetProperty(name string, defaultValue string)
	Clear()

	Getter() PropertyGetter
	Setter() PropertySetter

	Export(map[string]string) map[string]string
	Import(map[string]string)
}

// PropertyGetter 接口用来获取 Properties 中的属性
type PropertyGetter interface {
	Error() error
	SetError(err error)
	CleanError()

	GetString(name string, defaultValue string) string
	GetBool(name string, defaultValue bool) bool

	GetInt(name string, defaultValue int) int
	GetInt8(name string, defaultValue int8) int8
	GetInt16(name string, defaultValue int16) int16
	GetInt32(name string, defaultValue int32) int32
	GetInt64(name string, defaultValue int64) int64

	GetUint(name string, defaultValue uint) uint
	GetUint8(name string, defaultValue uint8) uint8
	GetUint16(name string, defaultValue uint16) uint16
	GetUint32(name string, defaultValue uint32) uint32
	GetUint64(name string, defaultValue uint64) uint64

	GetFloat32(name string, defaultValue float32) float32
	GetFloat64(name string, defaultValue float64) float64
}

// PropertySetter 接口用来设置 Properties 中的属性
type PropertySetter interface {
	SetString(name string, defaultValue string)
	SetBool(name string, defaultValue bool)

	SetInt(name string, defaultValue int)
	SetInt8(name string, defaultValue int8)
	SetInt16(name string, defaultValue int16)
	SetInt32(name string, defaultValue int32)
	SetInt64(name string, defaultValue int64)

	SetUint(name string, defaultValue uint)
	SetUint8(name string, defaultValue uint8)
	SetUint16(name string, defaultValue uint16)
	SetUint32(name string, defaultValue uint32)
	SetUint64(name string, defaultValue uint64)

	SetFloat32(name string, defaultValue float32)
	SetFloat64(name string, defaultValue float64)
}
