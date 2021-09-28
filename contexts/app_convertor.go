package contexts

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

// ForApplicationContext 为上下文创建转换器
func ForApplicationContext(ctx application.Context) Convertor {
	return &convertorForApp{ctx: ctx}
}

type convertorForApp struct {
	ctx application.Context
}

func (inst *convertorForApp) ToApplication() (application.Context, error) {
	return inst.ctx, nil
}

func (inst *convertorForApp) ToLang() (lang.Context, error) {

	holder := getHolderForAC(inst.ctx)

	lc := holder.lc

	if lc == nil {

	}

	return lc, nil
}
