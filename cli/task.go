package cli

type Task interface {
	Execute() error
	ExecuteWithArguments(args []string) error
}
