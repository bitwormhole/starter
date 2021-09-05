package tests

import (
	"errors"
	"os"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

type testPropertiesInGitLoader struct {
}

func (inst *testPropertiesInGitLoader) load() (collection.Properties, error) {

	dotgit, err := inst.findDotGitDir()
	if err != nil {
		return nil, err
	}

	file := dotgit.GetChild("test.properties")
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}

	return collection.ParseProperties(text, nil)
}

func (inst *testPropertiesInGitLoader) findDotGitDir() (fs.Path, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	pwd := fs.Default().GetPath(path)
	p := pwd
	for timeout := 99; p != nil; timeout-- {
		if timeout < 0 {
			break
		}
		dotgit := p.GetChild(".git")
		if dotgit.IsDir() {
			return dotgit, nil
		}
		p = p.Parent()
	}
	return nil, errors.New("cannot find '.git' in path " + pwd.Path())
}
