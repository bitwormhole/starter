package configen

type componentBuildingInfo struct {
	id      string
	aliases []string
	classes []string

	typeFullName    string
	typeShortName   string
	typePackageName string
	typeImporterTag string

	initMethod    string
	destroyMethod string
	inject        string
	scope         string
}
