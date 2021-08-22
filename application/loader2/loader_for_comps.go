package loader2

import (
	"errors"
	"strconv"

	"github.com/bitwormhole/starter/application"
)

// 组件加载器
type componentsLoader struct {
	countComponent int
	loading        *contextLoading
}

func (inst *componentsLoader) init(loading *contextLoading) {
	inst.loading = loading
}

func (inst *componentsLoader) load() error {

	ctx := inst.loading.context
	cfg := inst.loading.config
	dst := ctx.GetComponents()
	src := cfg.GetComponents()
	holders := make(map[string]application.ComponentHolder)

	// load holders

	for _, info := range src {
		h, err := inst.prepareComponent(info)
		if err != nil {
			return err
		}
		err = inst.putComponentTo(holders, h)
		if err != nil {
			return err
		}
	}

	dst.Import(holders)
	dst.GroupManager().Reload()

	// inject & init
	ciLoader := &comInstanceLoader{}
	ciLoader.init(inst.loading.context, false)
	ciLoader.addInstances(inst.findAllSingletonComs())
	return ciLoader.Load()
}

func (inst *componentsLoader) findAllSingletonComs() []application.ComponentInstance {
	dst := make([]application.ComponentInstance, 0, 16)
	comps := inst.loading.context.GetComponents()
	all := comps.Export(nil)
	for _, h := range all {
		info := h.GetInfo()
		if info.GetScope() == application.ScopeSingleton {
			dst = append(dst, h.GetInstance())
		}
	}
	return dst
}

func (inst *componentsLoader) putComponentTo(table map[string]application.ComponentHolder, h application.ComponentHolder) error {

	info := h.GetInfo()
	id := info.GetID()
	aliases := info.GetAliases()

	// make keys
	keys := make(map[string]string)
	for _, key := range aliases {
		keys[key] = key
	}
	keys[id] = id

	// put com(s) to table
	for _, name := range keys {
		older := table[name]
		if older != nil {
			return errors.New("The component.id(alias) is duplicated: " + name)
		}
		table[name] = h
	}
	return nil
}

func (inst *componentsLoader) prepareComponent(info application.ComponentInfo) (application.ComponentHolder, error) {

	info2, err := inst.prepareComInfo(info)
	if err != nil {
		return nil, err
	}

	h, err := inst.prepareComHolder(info2)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (inst *componentsLoader) prepareComHolder(info application.ComponentInfo) (application.ComponentHolder, error) {

	var h application.ComponentHolder = nil
	scope := info.GetScope()
	ctx := inst.loading.context
	switch scope {
	case application.ScopeSingleton:
		h = (&singletonComponentHolder{}).init(ctx, nil, info)
		break
	case application.ScopePrototype:
		h = (&prototypeComponentHolder{}).init(ctx, info)
		break
	case application.ScopeContext:
		break
	default:
		break
	}
	if h == nil {
		n := int(scope)
		return nil, errors.New("Unsupported scope:" + strconv.Itoa(n))
	}
	return h, nil
}

func (inst *componentsLoader) prepareComInfo(info application.ComponentInfo) (application.ComponentInfo, error) {

	inst.countComponent++
	n := inst.countComponent

	builder := &runtimeComInfoBuilder{}
	builder.init(info)
	builder.setDefaultID("__unnamed_component_" + strconv.Itoa(n))
	return builder.Create()
}

////////////////////////////////////////////////////////////////////////////////
