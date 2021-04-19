package collection

type Parameters interface {
	GetParam(name string) (string, error)

	Import(map[string]string)
	Export(map[string]string) map[string]string
}
