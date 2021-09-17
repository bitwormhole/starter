package cli

type Service interface {
	FindHandler(name string) (Handler, error)
}
