package collection

type SimpleParameters struct {
}

func (inst *SimpleParameters) GetParam(name string) (string, error) {
	// todo
	return "", nil
}

func (inst *SimpleParameters) Export(table map[string]string) map[string]string {
	// todo
	return nil
}

func (inst *SimpleParameters) Import(src map[string]string) {
	// todo

}

func CreateParameters() Parameters {
	return &SimpleParameters{}
}
