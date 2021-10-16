package cli

import "github.com/bitwormhole/starter/markup"

type Context struct {
	markup.Component `id:"cli-context"`

	Service       Service       `inject:"#cli-service"`
	ClientFactory ClientFactory `inject:"#cli-client-factory"`
	Filters       []Filter      `inject:".cli-filter"`
	Handlers      []Handler     `inject:".cli-handler"`
}
