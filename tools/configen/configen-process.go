package configen

import (
	"fmt"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

type configenProcess struct {
	context *configenContext
}

func (proc *configenProcess) run() error {

	tc := &lang.TryChain{}

	tc.Try(func() error {
		return proc.init()

	}).Try(func() error {
		return proc.loadInputFile()

	}).Try(func() error {
		return proc.loadComponents()

	}).Try(func() error {
		return proc.makeComBuildingInfoList()

	}).Try(func() error {
		return proc.buildCode()

	}).Try(func() error {
		return proc.saveCodeToOutputFile()

	}).Try(func() error {
		return nil
	})

	return tc.Result()
}

func (inst *configenProcess) init() error {

	ctx := inst.context
	pwd := ctx.pwd

	ctx.outputFileName = ctx.inputFileName + "-auto-gen.go"
	ctx.inputFile = pwd.GetChild(ctx.inputFileName)
	ctx.outputFile = pwd.GetChild(ctx.outputFileName)
	ctx.singleConfigFiles = map[string]*singleConfigFile{}
	ctx.properties = collection.CreateProperties()
	ctx.importerTable = map[string]*importerBuildingInfo{}
	ctx.importerManager = &importerManager{context: ctx}

	// 引入基本的包
	ctx.importerManager.loadImporter("github.com/bitwormhole/starter/lang")
	ctx.importerManager.loadImporter("github.com/bitwormhole/starter/application")
	ctx.importerManager.loadImporter("github.com/bitwormhole/starter/application/config")

	return nil
}

func (inst *configenProcess) loadInputFile() error {
	loader := &singleConfigFileLoader{context: inst.context}
	return loader.loadAllSingleFiles()
}

func (inst *configenProcess) loadComponents() error {
	loader := &componentDescriptorLoader{context: inst.context}
	return loader.load()
}

func (inst *configenProcess) makeComBuildingInfoList() error {
	maker := &componentBuildingInfoMaker{context: inst.context}
	return maker.make()
}

func (inst *configenProcess) buildCode() error {
	gen := &golangCodeGenerator{context: inst.context}
	code, err := gen.generate()
	if err != nil {
		return err
	}
	inst.context.code = code
	return nil
}

func (inst *configenProcess) saveCodeToOutputFile() error {

	code := inst.context.code

	fmt.Println("output.code:")
	fmt.Println(code)

	// todo ...
	return inst.context.outputFile.GetIO().WriteText(code, nil)

}
