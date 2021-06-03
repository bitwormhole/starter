package config

type ComInfoBuilder struct {
	id      string
	class   string
	aliases string
	scope   string
}

func (inst *ComInfoBuilder) ID(id string) *ComInfoBuilder {
	inst.id = id
	return inst
}

func (inst *ComInfoBuilder) Class(cls string) *ComInfoBuilder {
	inst.class = cls
	return inst
}

func (inst *ComInfoBuilder) Aliases(aliases string) *ComInfoBuilder {
	inst.aliases = aliases
	return inst
}

func (inst *ComInfoBuilder) Scope(scope string) *ComInfoBuilder {
	inst.scope = scope
	return inst
}

func (inst *ComInfoBuilder) Create() *ComInfo {

	info := &ComInfo{}
	info.ID = inst.id
	info.Class = inst.class
	info.Scope = 0
	info.Aliases = nil

	// todo ...

	return info
}
