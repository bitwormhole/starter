package terminal

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/lang"
)

// Context 是终端（terminal）的上下文
type Context struct {
	client  cli.Client
	console cli.Console
	app     application.Context
	ctx     lang.Context
	prompt  string
	exit    bool
}
