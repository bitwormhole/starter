package contexts

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

type Convertor interface {
	ToApplication() (application.Context, error)
	ToLang() (lang.Context, error)
}
