package configen

import (
	"errors"
	"strings"

	lang "github.com/bitwormhole/starter/lang"
)

type golangCodeGenerator struct {
	context *configenContext

	markStr   string
	markNL    string
	markTab   string
	markSpace string
}

func (inst *golangCodeGenerator) initMarks() {
	inst.markNL = "\r\n"
	inst.markSpace = " "
	inst.markStr = "\""
	inst.markTab = "\t"
}

func (inst *golangCodeGenerator) generate() (string, error) {

	inst.initMarks()
	builder := &strings.Builder{}
	tc := &lang.TryChain{}

	tc.Try(func() error {
		return inst.buildPackage(builder)

	}).Try(func() error {
		return inst.buildImportList(builder)

	}).Try(func() error {
		return inst.buildConfigFuncBegin(builder)

	}).Try(func() error {
		return inst.buildConfigFuncBody(builder)

	}).Try(func() error {
		return inst.buildConfigFuncEnd(builder)

	}).Try(func() error {
		return nil

	}).Try(func() error {
		return nil
	})

	builder.WriteString(inst.markNL)
	err := tc.Result()
	code := builder.String()
	return code, err
}

func (inst *golangCodeGenerator) buildPackage(builder *strings.Builder) error {
	builder.WriteString("package ")
	builder.WriteString(inst.context.head.packageName)
	builder.WriteString(inst.markNL)

	// comment for file
	builder.WriteString("// This file is auto-generate by configen, never edit it.")
	builder.WriteString(inst.markNL)

	return nil
}

func (inst *golangCodeGenerator) buildImportList(builder *strings.Builder) error {

	mk_nl := inst.markNL
	mk_str := inst.markStr
	mk_sp := inst.markSpace
	items := inst.context.importerTable

	builder.WriteString(mk_nl)
	builder.WriteString("import (")
	builder.WriteString(mk_nl)

	for key := range items {
		item := items[key]
		builder.WriteString("\t")
		builder.WriteString(item.tag + mk_sp)
		builder.WriteString(mk_str + item.fullName + mk_str + mk_nl)
	}

	builder.WriteString(")" + mk_nl)
	return nil
}

func (inst *golangCodeGenerator) buildConfigFuncBegin(builder *strings.Builder) error {

	mk_nl := inst.markNL

	// func Config(){
	cfg_fn_name := inst.context.head.configFunctionName
	builder.WriteString(mk_nl + "func ")
	builder.WriteString(cfg_fn_name)
	builder.WriteString("(cfg application.ConfigBuilder){" + mk_nl)
	return nil
}

func (inst *golangCodeGenerator) buildConfigFuncEnd(builder *strings.Builder) error {
	// }
	builder.WriteString("}" + inst.markNL)
	return nil
}

func (inst *golangCodeGenerator) buildConfigFuncBody(builder *strings.Builder) error {
	list := inst.context.comBuildingInfoList
	for index := range list {
		item := list[index]
		err := inst.doBuildConfigForComponent(item, builder)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *golangCodeGenerator) doBuildConfigForComponent(com *componentBuildingInfo, builder *strings.Builder) error {

	tc := &lang.TryChain{}
	mk_nl := inst.markNL
	mk_str := inst.markStr
	mk_tab := inst.markTab

	importer, err := inst.context.importerManager.getImporter(com.typePackageName)
	if err != nil {
		return err
	}

	tc.Try(func() error {
		// 注释
		builder.WriteString(mk_nl)
		builder.WriteString(mk_tab + "// " + com.id + mk_nl)
		return nil
	}).Try(func() error {
		// com begin
		builder.WriteString("\tcfg.AddComponent(&config.ComInfo{")
		builder.WriteString(mk_nl)
		return nil
	}).Try(func() error {
		// id
		id := com.id
		if id == "" {
			return nil
		}
		builder.WriteString("\t\tID:" + mk_str)
		builder.WriteString(id)
		builder.WriteString(mk_str + "," + mk_nl)
		return nil
	}).Try(func() error {
		// class
		classlist := com.classes
		if classlist == nil {
			return nil
		}
		if len(classlist) < 1 {
			return nil
		}
		sep := ""
		builder.WriteString("\t\tClass:" + mk_str)
		for index := range classlist {
			cls := classlist[index]
			builder.WriteString(sep)
			builder.WriteString(cls)
			sep = " "
		}
		builder.WriteString(mk_str + "," + mk_nl)
		return nil
	}).Try(func() error {
		// aliases
		aliases := com.aliases
		if aliases == nil {
			return nil
		}
		if len(aliases) < 1 {
			return nil
		}
		sep := ""
		builder.WriteString(mk_tab + mk_tab + "Aliases:[]string{")
		for index := range aliases {
			alias := aliases[index]
			builder.WriteString(sep)
			builder.WriteString(mk_str + alias + mk_str)
			sep = ","
		}
		builder.WriteString("}," + mk_nl)
		return nil

	}).Try(func() error {
		// scope
		scope := com.scope
		token := ""
		if scope == "" {
			return nil
		} else if scope == "singleton" {
			token = "application.ScopeSingleton"
		} else if scope == "prototype" {
			token = "application.ScopePrototype"
		} else {
			return errors.New("bad scope name:" + scope)
		}
		builder.WriteString(mk_tab + mk_tab + "Scope: ")
		builder.WriteString(token)
		builder.WriteString("," + mk_nl)
		return nil
	}).Try(func() error {
		return nil
	}).Try(func() error {
		return nil

	}).Try(func() error {
		// OnInject
		method := com.inject
		if method == "" {
			return nil
		}
		tag := com.typeImporterTag + "." + com.typeShortName

		builder.WriteString(mk_tab + mk_tab)
		builder.WriteString("OnInject: func(obj lang.Object, context application.RuntimeContext) error {" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + mk_tab)
		builder.WriteString("t := obj.(*" + tag + ")" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + mk_tab)
		builder.WriteString("return " + method + "(t,context)" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + "}," + mk_nl)
		return nil
	}).Try(func() error {
		// OnInit
		method := com.initMethod
		if method == "" {
			return nil
		}
		tag := com.typeImporterTag + "." + com.typeShortName

		builder.WriteString(mk_tab + mk_tab)
		builder.WriteString("OnInit: func(obj lang.Object) error {" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + mk_tab)
		builder.WriteString("t := obj.(*" + tag + ")" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + mk_tab)
		builder.WriteString("return t." + method + "()" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + "}," + mk_nl)
		return nil
	}).Try(func() error {
		// OnDestroy
		method := com.destroyMethod
		if method == "" {
			return nil
		}
		tag := com.typeImporterTag + "." + com.typeShortName

		builder.WriteString(mk_tab + mk_tab)
		builder.WriteString("OnDestroy: func(obj lang.Object) error {" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + mk_tab)
		builder.WriteString("t := obj.(*" + tag + ")" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + mk_tab)
		builder.WriteString("return t." + method + "()" + mk_nl)

		builder.WriteString(mk_tab + mk_tab + "}," + mk_nl)
		return nil
	}).Try(func() error {
		// OnNew
		tag := importer.tag + "." + com.typeShortName
		builder.WriteString(mk_tab + mk_tab + "OnNew: func() lang.Object {" + mk_nl)
		builder.WriteString(mk_tab + mk_tab + mk_tab + "return &" + tag + "{}" + mk_nl)
		builder.WriteString(mk_tab + mk_tab + "}," + mk_nl)
		return nil
	}).Try(func() error {
		// com end
		builder.WriteString(mk_tab + "})" + mk_nl)
		return nil
	})

	return tc.Result()
}
