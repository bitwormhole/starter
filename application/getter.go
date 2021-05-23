package application

// ContextGetter 接口向 Context 的使用者提供简易的 getter 方法
type ContextGetter interface {

	// for property
	GetProperty(name string) (string, error)
	GetPropertySafely(name string, _default string) string
	GetPropertyString(name string, _default string) string
	GetPropertyInt(name string, _default int) int
}
