package runtime

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// RuntimeContextLoader 运行时上下文加载器
type RuntimeContextLoader struct {
	comInfoList []application.ComponentInfo
	comTable    map[string]application.ComponentHolder
	context     application.RuntimeContext
	config      application.Configuration
	args        []string
}

// Load 方法根据传入的配置加载运行时上下文
func (inst *RuntimeContextLoader) Load(config application.Configuration, args []string) (application.RuntimeContext, error) {

	inst.config = config
	inst.comTable = make(map[string]application.ComponentHolder)
	inst.comInfoList = nil
	inst.context = nil
	inst.args = args

	tc := &lang.TryChain{}

	tc.Try(func() error {
		return inst.createRuntimeContext()

	}).Try(func() error {
		return inst.loadArguments()

	}).Try(func() error {
		return inst.loadPropertiesInRes1()

	}).Try(func() error {
		return inst.loadPropertiesInRes2()

	}).Try(func() error {
		return inst.prepareComInfoList()

	}).Try(func() error {
		return inst.doCreateComponents()

	}).Try(func() error {
		return inst.loadSingletonComponents()

	}).Try(func() error {
		return nil

	}).Try(func() error {
		return inst.logDebugInfo()
	})

	err := tc.Result()
	ctx := inst.context
	if err != nil {
		ctx = nil
	}
	return ctx, err
}

func (inst *RuntimeContextLoader) loadArguments() error {
	src := inst.args
	dst := inst.context.GetArguments()
	if src == nil {
		return nil
	}
	dst.Import(src)
	return nil
}

func (inst *RuntimeContextLoader) loadPropertiesInArgs() error {

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
	key := "bitwormhole.profiles.active"
	properties := inst.context.GetProperties()
	profile := properties.GetProperty(key, "default")
	name := "/application-" + profile + ".properties"
	return inst.loadPropertiesInRes(name)
}

func (inst *RuntimeContextLoader) createRuntimeContext() error {

	core := createRuntimeContextCore()

	core.appName = ""
	core.appVersion = ""
	core.time1 = 0
	core.time2 = 0
	core.uri = ""
	core.resources = inst.config.GetResources()

	inst.context = core.context
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

	scopeWant := application.ScopeSingleton

	cc := inst.context.OpenCreationContext(scopeWant)
	context := cc.GetContext()
	components := context.GetComponents()
	table := components.Export(nil)

	for name := range table {
		holder := table[name]
		info := holder.GetInfo()
		id := info.GetID()
		scope := info.GetScope()
		if (id == name) && (scope == scopeWant) {
			_, err := components.GetComponent(name)
			if err != nil {
				return err
			}
		}
	}

	return cc.Close()
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
