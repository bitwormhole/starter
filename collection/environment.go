package collection

type Environment interface {
	GetEnv(name string) (string, error)
	SetEnv(name string, value string)

	Import(map[string]string)
	Export(map[string]string) map[string]string
}
