package configen

type componentDescriptor struct {
	descriptorName string
	typeName       string

	id            string
	class         string
	aliases       string
	initMethod    string
	destroyMethod string
	inject        string
	scope         string
}
