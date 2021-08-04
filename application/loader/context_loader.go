package loader

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/runtime"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

// RuntimeContextLoader 运行时上下文加载器
type RuntimeContextLoader struct {
	comInfoList []application.ComponentInfo
	comTable    map[string]application.ComponentHolder
	context     application.Context
	config      application.Configuration
	args        []string
}

// Load 方法根据传入的配置加载运行时上下文
func (inst *RuntimeContextLoader) Load(config application.Configuration, args []string) (application.Context, error) {

	inst.config = config
	inst.comTable = make(map[string]application.ComponentHolder)
	inst.comInfoList = nil
	inst.context = nil
	inst.args = args

	err := inst.createRuntimeContext()
	if err != nil {
		return nil, err
	}

	err = inst.loadArguments()
	if err != nil {
		return nil, err
	}

	err = inst.loadEnv()
	if err != nil {
		return nil, err
	}

	err = inst.loadDefaultProperties()
	if err != nil {
		return nil, err
	}

	err = inst.loadPropertiesInArgs()
	if err != nil {
		return nil, err
	}

	err = inst.loadPropertiesInRes1()
	if err != nil {
		return nil, err
	}

	err = inst.loadPropertiesInRes2()
	if err != nil {
		return nil, err
	}

	err = inst.loadPropertiesInLocalFile()
	if err != nil {
		return nil, err
	}

	err = inst.loadAtts()
	if err != nil {
		return nil, err
	}

	err = inst.prepareComInfoList()
	if err != nil {
		return nil, err
	}

	err = inst.doCreateComponents()
	if err != nil {
		return nil, err
	}

	err = inst.loadSingletonComponents()
	if err != nil {
		return nil, err
	}

	// return inst.logDebugInfo()

	ctx := inst.context
	return ctx, nil
}

func (inst *RuntimeContextLoader) loadArguments() error {
	src := inst.args
	dest := inst.context.GetArguments()
	if src == nil {
		return nil
	}
	dest.Import(src)
	return nil
}

func (inst *RuntimeContextLoader) loadEnv() error {
	src := inst.config.GetEnvironment()
	dst := inst.context.GetEnvironment()
	table := make(map[string]string)
	if src != nil {
		table = src.Export(table)
	} else {
		array := os.Environ()
		for index := range array {
			item := array[index]
			idx := strings.Index(item, "=")
			if idx < 0 {
				continue
			}
			key := strings.TrimSpace(item[0:idx])
			val := strings.TrimSpace(item[idx+1:])
			table[key] = val
		}
	}
	dst.Import(table)
	return nil
}

func (inst *RuntimeContextLoader) loadDefaultProperties() error {
	src := inst.config.GetDefaultProperties()
	dst := inst.context.GetProperties()
	table := src.Export(nil)
	for key := range table {
		val := table[key]
		dst.SetProperty(key, val)
	}
	return nil
}

func (inst *RuntimeContextLoader) loadPropertiesInArgs() error {

	enable := inst.config.IsEnableLoadPropertiesFromArguments()
	if !enable {
		// skip
		return nil
	}

	props := inst.context.GetProperties()
	args := inst.context.GetArguments()
	array := args.Export()

	//	fmt.Println(props.GetProperty("", "args.array:"))

	for index := range array {
		text := array[index]
		//	fmt.Println("   [args.item text:", text, "]")
		if !strings.HasPrefix(text, "--") {
			continue
		}
		text = strings.TrimLeft(text, "-")
		parts := strings.SplitN(text, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			val := parts[1]
			props.SetProperty(key, val)
			//	fmt.Println("   [args.item key:", key, " value:", val, "]")
		}
	}

	return nil
}

func (inst *RuntimeContextLoader) loadPropertiesInRes(resourceName string) error {
	text, err := inst.context.GetResources().GetText(resourceName)
	if err != nil {
		return nil
	}
	properties := inst.context.GetProperties()
	properties, err = collection.ParseProperties(text, properties)
	if err != nil {
		return err
	}
	return inst.loadPropertiesInArgs()
}

func (inst *RuntimeContextLoader) loadPropertiesInRes1() error {
	name := "/application.properties"
	return inst.loadPropertiesInRes(name)
}

func (inst *RuntimeContextLoader) loadPropertiesInRes2() error {
	key := "application.profiles.active"
	properties := inst.context.GetProperties()
	profile := properties.GetProperty(key, "default")
	name := "/application-" + profile + ".properties"
	log.Println(key+":", profile)
	return inst.loadPropertiesInRes(name)
}

func (inst *RuntimeContextLoader) loadAtts() error {
	src := inst.config.GetAttributes()
	dst := inst.context.GetAttributes()
	if src == nil || dst == nil {
		return errors.New("ptr==nil")
	}
	table := src.Export(nil)
	dst.Import(table)
	return nil
}

func (inst *RuntimeContextLoader) loadPropertiesInLocalFile() error {
	const undef = ""
	const key = "application.properties"
	path := inst.context.GetProperties().GetProperty(key, undef)
	if path == undef {
		return nil
	}
	// read text
	file := fs.Default().GetPath(path)
	text, err := file.GetIO().ReadText()
	if err != nil {
		return err
	}
	// load properties
	properties := inst.context.GetProperties()
	properties, err = collection.ParseProperties(text, properties)
	if err != nil {
		return err
	}
	return inst.loadPropertiesInArgs()
}

func (inst *RuntimeContextLoader) createRuntimeContext() error {

	builder := &runtime.RuntimeContextBuilder{}

	builder.AppName = ""
	builder.AppVersion = ""
	builder.Time1 = 0
	builder.Time2 = 0
	builder.URL = ""
	builder.Resources = inst.config.GetResources()
	builder.ComLoader = &StandardComponentLoader{}

	context, err := builder.Create()
	if err != nil {
		return err
	}
	inst.context = context
	return nil
}

func (inst *RuntimeContextLoader) prepareComInfoList() error {
	src := inst.config.GetComponents()
	dst := make([]application.ComponentInfo, 0)
	preprocessor := &componentInfoPreprocessor{}
	for index := range src {
		info := src[index]
		info, err := preprocessor.prepare(info, index)
		if err != nil {
			return err
		}
		dst = append(dst, info)
	}
	inst.comInfoList = dst
	return nil
}

func (inst *RuntimeContextLoader) doCreateComponents() error {

	// 根据 info 创建 对应的 holder

	ctx := inst.context
	src := inst.comInfoList
	dst := make(map[string]application.ComponentHolder)

	for index := range src {
		info := src[index]
		scope := info.GetScope()
		var holder application.ComponentHolder
		if scope == application.ScopeSingleton {
			holder = &SingletonComponentHolder{context: ctx, info: info}
		} else if scope == application.ScopePrototype {
			holder = &PrototypeComponentHolder{context: ctx, info: info}
		} else if scope == application.ScopeContext {
			continue
		} else {
			continue
		}
		err := inst.putComHolderToTable(dst, holder)
		if err != nil {
			return err
		}
	}

	// 导入到 context 里
	com_set := ctx.GetComponents()
	com_set.Import(dst)
	inst.comTable = com_set.Export(nil)

	return nil
}

func (inst *RuntimeContextLoader) putComHolderToTable(table map[string]application.ComponentHolder, holder application.ComponentHolder) error {

	info := holder.GetInfo()
	id := info.GetID()
	aliases := info.GetAliases()

	id_in_list := false
	for index := range aliases {
		name := aliases[index]
		if name == id {
			id_in_list = true
			break
		}
	}
	if !id_in_list {
		aliases = append(aliases, id)
	}

	for index := range aliases {
		name := aliases[index]
		older := table[name]
		if older != nil {
			return errors.New("the ID (alias) of component is duplicate:" + name)
		}
		table[name] = holder
	}

	return nil
}

func (inst *RuntimeContextLoader) loadSingletonComponents() error {

	context := inst.context
	components := context.GetComponents()
	table := components.Export(nil)
	injector := inst.context.Injector()
	injection, err := injector.OpenInjection(context)

	if err != nil {
		return err
	}

	for name := range table {
		holder := table[name]
		info := holder.GetInfo()
		id := info.GetID()
		scope := info.GetScope()
		if (id == name) && (scope == application.ScopeSingleton) {
			src := injection.Select("#" + name)
			_, err := src.Read()
			if err != nil {
				return err
			}
		}
	}

	pool1 := injection.Pool()
	pool2 := context.GetReleasePool()
	ppbinding := &runtime.PoolPairBinding{}
	ppbinding.Init(pool1, pool2)

	return injection.Close()
}

func (inst *RuntimeContextLoader) logDebugInfo() error {

	props := inst.context.GetProperties()
	table := props.Export(nil)
	keys := make([]string, 0)

	for key := range table {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	fmt.Println("context.properties:")

	for index := range keys {
		k := keys[index]
		v := table[k]
		fmt.Println("  " + k + "=[" + v + "]")
	}

	return nil
}
