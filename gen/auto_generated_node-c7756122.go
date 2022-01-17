// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	bootstrap0x1b594d "github.com/bitwormhole/starter/bootstrap"
	lang0xbf4f1f "github.com/bitwormhole/starter/lang"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComBoot struct {
	instance *bootstrap0x1b594d.Boot
	 markup0x23084a.Component `id:"main-looper"`
	Lives []lang0xbf4f1f.Object `inject:".life"`
	Concurrent bool `inject:"${application.loopers.concurrent}"`
}

