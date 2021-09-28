package contexts

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

const contextHolderBindingName = "/bitwormhole/starter/contexts/holder#binding"

type contextHolder struct {
	lc lang.Context
	ac application.Context
}

func getHolderForLC(ctx lang.Context) *contextHolder {
	const name = contextHolderBindingName
	o1 := ctx.GetValue(name)
	if o1 != nil {
		o2, ok := o1.(*contextHolder)
		if ok {
			return o2
		}
	}
	holder := &contextHolder{}
	ctx.SetValue(name, holder)
	holder.lc = ctx
	return holder
}

func getHolderForAC(ctx application.Context) *contextHolder {
	const name = contextHolderBindingName
	atts := ctx.GetAttributes()
	o1, err := atts.GetAttributeRequired(name)
	if err == nil {
		o2, ok := o1.(*contextHolder)
		if ok {
			return o2
		}
	}
	holder := &contextHolder{}
	atts.SetAttribute(name, holder)
	holder.ac = ctx
	return holder
}

////////////////////////////////////////////////////////////////////////////////

// SetupLC 为上下文设置绑定
func SetupLC(ctx lang.Context) lang.Context {
	if ctx == nil {
		return nil
	}
	holder := getHolderForLC(ctx)
	holder.lc = ctx
	return ctx
}

// SetupAC 为上下文设置绑定
func SetupAC(ctx application.Context) application.Context {
	if ctx == nil {
		return nil
	}
	holder := getHolderForAC(ctx)
	holder.ac = ctx
	return ctx
}
