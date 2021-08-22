package loader2

import (
	"errors"

	"github.com/bitwormhole/starter/application"
)

type comGroupsBuilder struct {
	holders map[string]application.ComponentHolder
	groups  map[string]*componentGroup
}

func (inst *comGroupsBuilder) init() {
	inst.groups = make(map[string]*componentGroup)
	inst.holders = make(map[string]application.ComponentHolder)
}

func (inst *comGroupsBuilder) getGroup(selector string, create bool) (*componentGroup, error) {
	g := inst.groups[selector]
	if g == nil {
		if create {
			g = &componentGroup{}
			g.init(selector)
			inst.groups[selector] = g
		}
	}
	if g == nil {
		return nil, errors.New("no group with selector:" + selector)
	}
	return g, nil
}

func (inst *comGroupsBuilder) add(name string, holder application.ComponentHolder) {

	info := holder.GetInfo()
	id := info.GetID()
	classes := info.GetClasses()
	aliases := info.GetAliases()

	older := inst.holders[id]
	if older == nil {
		inst.holders[id] = holder
	} else {
		return
	}

	for _, alias := range aliases {
		sel := "#" + alias
		g, _ := inst.getGroup(sel, true)
		g.add(holder)
	}

	for _, cname := range classes {
		sel := "." + cname
		g, _ := inst.getGroup(sel, true)
		g.add(holder)
	}

	// with id
	g, _ := inst.getGroup("#"+id, true)
	g.add(holder)

	// with *
	g, _ = inst.getGroup("*", true)
	g.add(holder)
}

func (inst *comGroupsBuilder) createTo(holders map[string]application.ComponentHolder, groups map[string]application.ComponentGroup) {

	src1 := inst.groups
	src2 := inst.holders
	dst1 := groups
	dst2 := holders

	for key, val := range src1 {
		dst1[key] = val
	}

	for key, val := range src2 {
		dst2[key] = val
	}

}
