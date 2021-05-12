package configen2

import "github.com/bitwormhole/starter/io/fs"

type configen2context struct {
	inputFileName  string
	outputFileName string
	inputFile      fs.Path
	outputFile     fs.Path
	pwd            fs.Path
	code           string
	com            *codeObjectModel
}
