package collection

type SimpleArguments struct {
	args []string
}

func (inst *SimpleArguments) GetArgument(name string) (string, error) {
	// todo
	return "", nil
}

func (inst *SimpleArguments) Export() []string {
	return inst.args
}

func (inst *SimpleArguments) Import(args []string) {
	inst.args = args
}

func CreateArguments() Arguments {
	return &SimpleArguments{}
}
