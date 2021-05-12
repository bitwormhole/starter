package configen2

import (
	"encoding/json"
	"fmt"
)

type configen2process struct {
	context *configen2context
}

func (inst *configen2process) run() error {

	err := inst.init()
	if err != nil {
		return err
	}

	err = inst.loadInputFile() // the starter.config
	if err != nil {
		return err
	}

	err = inst.scanGolangSourceFiles()
	if err != nil {
		return err
	}

	err = inst.saveToOutputFile()
	if err != nil {
		return err
	}

	return nil
}

func (inst *configen2process) init() error {

	ctx := inst.context

	ctx.inputFileName = "starter.config"
	ctx.outputFileName = "starter.config-auto-gen.go"

	ctx.inputFile = ctx.pwd.GetChild(ctx.inputFileName)
	ctx.outputFile = ctx.pwd.GetChild(ctx.outputFileName)

	return nil
}

func (inst *configen2process) loadInputFile() error {

	file := inst.context.inputFile
	fmt.Println("load config from ", file.Path())

	return nil
}

func (inst *configen2process) saveToOutputFile() error {

	file := inst.context.outputFile
	fmt.Println("write to ", file.Path())

	return nil
}

func (inst *configen2process) scanGolangSourceFiles() error {
	scanner := &golangSourceScanner{context: inst.context}
	err := scanner.scan()
	if err != nil {
		return err
	}

	// print COM
	com := inst.context.com
	js, err := json.Marshal(com)
	if err != nil {
		return err
	}
	fmt.Println("COM(Code Object Model):", js)

	return nil
}
