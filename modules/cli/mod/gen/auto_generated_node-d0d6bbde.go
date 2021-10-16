// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	context0xec2727 "context"
	cli0xfc9cfc "github.com/bitwormhole/starter/cli"
	filters0xe74833 "github.com/bitwormhole/starter/cli/filters"
	support0x409e86 "github.com/bitwormhole/starter/cli/support"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComContext struct {
	instance *cli0xfc9cfc.Context
	 markup0x23084a.Component `id:"cli-context"`
	Service cli0xfc9cfc.Service `inject:"#cli-service"`
	ClientFactory cli0xfc9cfc.ClientFactory `inject:"#cli-client-factory"`
	Filters []cli0xfc9cfc.Filter `inject:".cli-filter"`
	Handlers []cli0xfc9cfc.Handler `inject:".cli-handler"`
}


type pComHandlerFinderFilter struct {
	instance *filters0xe74833.HandlerFinderFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"800"`
}


type pComContextFilter struct {
	instance *filters0xe74833.ContextFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"900"`
	Context context0xec2727.Context `inject:"context"`
	Service cli0xfc9cfc.Service ``
}


type pComExecutorFilter struct {
	instance *filters0xe74833.ExecutorFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"700"`
}


type pComMultilineSupportFilter struct {
	instance *filters0xe74833.MultilineSupportFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"850"`
}


type pComNopFilter struct {
	instance *filters0xe74833.NopFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"0"`
}


type pComDefaultClientFactory struct {
	instance *support0x409e86.DefaultClientFactory
	 markup0x23084a.Component `id:"cli-client-factory"`
	CLI *cli0xfc9cfc.Context `inject:"#cli-context"`
}


type pComDefaultSerivce struct {
	instance *support0x409e86.DefaultSerivce
	 markup0x23084a.Component `id:"cli-service" initMethod:"Init"`
	CLI *cli0xfc9cfc.Context `inject:"#cli-context"`
	handlerTable map[string]cli0xfc9cfc.Handler ``
	chain cli0xfc9cfc.FilterChain ``
	chainBuilder *cli0xfc9cfc.FilterChainBuilder ``
}

