package configen2

type codeObjectModel struct {
	injectors map[string]*comInjector
}

type comInjector struct {
	key         string
	shortName   string
	fullName    string
	packageName string
	packageTag  string

	context *comInjectorField
	target  *comInjectorField

	fields map[string]*comInjectorField
}

type comInjectorField struct {
	isTarget  bool
	isContext bool

	name     string
	typeName string
	rawTag   string
	tags     map[string]string
}
