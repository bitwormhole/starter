package application

import "github.com/bitwormhole/starter/lang"

type Injection interface {
	AsList() Injection
	IncludeAliases() Injection
	Accept(f ComponentHolderFilter) Injection
	To(func(lang.Object) bool)
}

type Injector interface {
	ContextGetter

	//  selector 类似css中的selector，用“#xxx”表示id选择器，用“.xxx”表示类选择器，用“*”表示全部
	Inject(selector string) Injection
	GetComponent(selector string) (lang.Object, error)

	Done() error
}
