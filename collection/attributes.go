package collection

import "github.com/bitwormhole/starter/lang"

type Attributes interface {
	GetAttribute(name string) (lang.Object, error)

	Import(map[string]lang.Object)
	Export(map[string]lang.Object) map[string]lang.Object
}
