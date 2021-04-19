package collection

// Properties 接口表示对属性列表的引用。
type Properties interface {
	GetPropertyRequired(name string) (string, error)
	GetProperty(name string, defaultValue string) string
	SetProperty(name string, value string)
	Clear()

	Export(map[string]string) map[string]string
	Import(map[string]string)
}
