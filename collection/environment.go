package collection

type Environment interface {
	GetEnv(name string) (string, error)

	Import(map[string]string)
	Export(map[string]string) map[string]string
}
