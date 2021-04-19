package collection

type SimpleEnvironment struct {
}

func (inst *SimpleEnvironment) GetEnv(name string) (string, error) {
	// todo
	return "", nil
}

func (inst *SimpleEnvironment) Export(dst map[string]string) map[string]string {
	// todo
	return nil
}

func (inst *SimpleEnvironment) Import(src map[string]string) {
	// todo

}

func CreateEnvironment() Environment {
	return &SimpleEnvironment{}
}
