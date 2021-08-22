package loader2

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/vlog"
)

type componentManager struct {
	groups  map[string]application.ComponentGroup  // map[selector] group
	all     map[string]application.ComponentHolder // map[id] holder
	context application.Context
}

func (inst *componentManager) _Impl() (application.Components, application.ComponentGroupManager) {
	return inst, inst
}

func (inst *componentManager) init(ctx application.Context) {
	inst.context = ctx
	inst.all = make(map[string]application.ComponentHolder)
	inst.groups = make(map[string]application.ComponentGroup)
}

///////////////////////////////
// impl ComponentGroupManager

func (inst *componentManager) GetGroup(selector string) application.ComponentGroup {
	return inst.groups[selector]
}

func (inst *componentManager) Reload() error {

	srclist := inst.all
	dstlist := make(map[string]application.ComponentGroup)
	builder := &comGroupsBuilder{}
	builder.init()

	for name, holder := range srclist {
		builder.add(name, holder)
	}

	builder.createTo(inst.all, dstlist)
	inst.groups = dstlist
	return nil
}

///////////////////////////////
// impl Components

func (inst *componentManager) GetComponentNameList(includeAliases bool) []string {
	srcSet := inst.all
	namelist := make([]string, 0, len(srcSet))
	for key, holder := range srcSet {
		info := holder.GetInfo()
		id := info.GetID()
		if key != id {
			continue // skip, if alias
		}
		namelist = append(namelist, id)
		if !includeAliases {
			continue // skip, if id only
		}
		aliases := info.GetAliases()
		for _, alias := range aliases {
			namelist = append(namelist, alias)
		}
	}
	return namelist
}

// getters
func (inst *componentManager) FindComponent(selector string) (application.ComponentHolder, error) {

	g := inst.groups[selector]
	if g == nil {
		return nil, errors.New("no component with selector: " + selector)
	}

	list := g.ListAll()
	const want = 1
	got := len(list)

	if got == want {
		return list[0], nil
	}

	builder := strings.Builder{}
	builder.WriteString("find component with selector:[")
	builder.WriteString(selector)
	builder.WriteString("] want:[")
	builder.WriteString(strconv.Itoa(want))
	builder.WriteString("] but got:[")
	builder.WriteString(strconv.Itoa(got))
	builder.WriteString("] item(s).")
	return nil, errors.New(builder.String())
}

func (inst *componentManager) FindComponents(selector string) []application.ComponentHolder {
	g := inst.groups[selector]
	if g == nil {
		logger := vlog.Default()
		if logger.IsDebugEnabled() {
			logger.Warn("no component with selector: " + selector)
		}
		return []application.ComponentHolder{}
	}
	return g.ListAll()
}

func (inst *componentManager) FindComponentsWithFilter(selector string, f application.ComponentHolderFilter) []application.ComponentHolder {
	src := inst.FindComponents(selector)
	if f == nil {
		return src
	}
	dst := make([]application.ComponentHolder, 0)
	for _, item := range src {
		name := item.GetInfo().GetID()
		if f(name, item) {
			dst = append(dst, item)
		}
	}
	return dst
}

func (inst *componentManager) GroupManager() application.ComponentGroupManager {
	return inst
}

// export & import
func (inst *componentManager) Export(dst map[string]application.ComponentHolder) map[string]application.ComponentHolder {
	if dst == nil {
		dst = make(map[string]application.ComponentHolder)
	}
	src := inst.all
	if src == nil {
		return dst
	}
	for id, item := range src {
		dst[id] = item
	}
	return dst
}

func (inst *componentManager) Import(src map[string]application.ComponentHolder) {
	if src == nil {
		return
	}
	dst := inst.all
	if dst == nil {
		dst = make(map[string]application.ComponentHolder)
	}
	ctx := inst.context
	for id, item := range src {
		child := item.MakeChild(ctx)
		dst[id] = child
	}
	inst.all = dst
}
