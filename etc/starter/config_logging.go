package starter

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/starter/vlog/std"
)

type theVlogDefaultContext struct {
	markup.Component
	instance *std.DefaultContext `id:"vlog-std-context"`

	DefaultLevel     string         `inject:"${vlog.default.level}"`
	DefaultFormatter vlog.Formatter `inject:"#vlog-default-formatter"`
	Channels         []vlog.Channel `inject:".vlog-std-channel"`
	MainChannel      vlog.Channel   `inject:"#vlog-std-main-channel"`
}

type theVlogLoggerFactory struct {
	markup.Component
	instance *std.StandardLoggerFactory `id:"vlog-std-logger-factory" initMethod:"Start" destroyMethod:"Stop"`

	Context std.Context `inject:"#vlog-std-context"`
}

type theVlogDefaultFormatter struct {
	markup.Component
	instance *std.DefaultFormatter `id:"vlog-default-formatter"`
}

type theVlogMainChannel struct {
	markup.Component
	instance *std.LogChannel `id:"vlog-std-main-channel" class:"vlog-std-channel"`

	Context   std.Context    `inject:"#vlog-std-context"`
	Name      string         `inject:"vlog-main"`
	Enable    bool           `inject:"${vlog.main.enable}"`
	Filters   []vlog.Filter  `x-inject:"*"`
	Writer    vlog.Writer    `inject:"#vlog-std-distributor"`
	Level     string         `inject:"${vlog.main.level}"`
	Formatter vlog.Formatter `x-inject:"*"`
}

type theVlogDistributor struct {
	markup.Component
	instance *std.Distributor `id:"vlog-std-distributor"`

	Channels []vlog.Channel `inject:".vlog-std-sub-channel"`
}

type theVlogConsoleChannel struct {
	markup.Component
	instance *std.LogChannel `id:"vlog-std-console-channel" class:"vlog-std-channel vlog-std-sub-channel"`

	Context   std.Context    `inject:"#vlog-std-context"`
	Name      string         `inject:"vlog-console"`
	Enable    bool           `inject:"${vlog.console.enable}"`
	Filters   []vlog.Filter  `x-inject:"*"`
	Writer    vlog.Writer    `inject:"#vlog-std-console-writer"`
	Level     string         `inject:"${vlog.console.level}"`
	Formatter vlog.Formatter `x-inject:"*"`
}

type theVlogConsoleWriter struct {
	markup.Component
	instance *std.ConsoleWriter `id:"vlog-std-console-writer"`
}
