package collection

type Arguments interface {
	//	Count() int
	//	GetArgumentAt(index int) (string, error)
	//	GetArgumentWithOffset(name string, offset int) (string, error)
	//	GetArgument(name string) (string, error)

	Import([]string)
	Export() []string
}
