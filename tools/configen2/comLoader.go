package configen2

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/bitwormhole/starter/io/fs"
)

type comLoader struct {
}

func (inst *comLoader) loadFromFile(golangSourceFile fs.Path) (*codeObjectModel, error) {
	text, err := golangSourceFile.GetIO().ReadText()
	if err != nil {
		return nil, err
	}
	return inst.loadSourceCode(text)
}

func (inst *comLoader) loadSourceCode(golangSourceCode string) (*codeObjectModel, error) {

	mode := parser.ParseComments
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", golangSourceCode, mode)
	if err != nil {
		return nil, err
	}

	comments := &commentSet{}
	comments.init(f.Comments)

	// builder := &comBuilder{}
	handler := DefaultHandler(comments)

	ast.Inspect(f, func(n ast.Node) bool {

		//  fmt.Println(n)
		// builder.hello()
		handler.handle(n)

		return true
	})

	return nil, nil
}
