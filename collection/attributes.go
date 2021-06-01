package collection

import "github.com/bitwormhole/starter/lang"

type Atts interface {
	GetAttribute(name string) lang.Object
	SetAttribute(name string, value lang.Object)
}

type Attributes interface {
	Atts

	GetAttributeRequired(name string) (lang.Object, error)

	Import(map[string]lang.Object)
	Export(map[string]lang.Object) map[string]lang.Object
}
