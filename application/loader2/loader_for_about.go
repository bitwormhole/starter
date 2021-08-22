package loader2

import "github.com/bitwormhole/starter/vlog"

type aboutInfoLoader struct{}

func (inst *aboutInfoLoader) load(loading *contextLoading) error {

	dp := loading.context.GetProperties()
	goVer := dp.GetProperty("go.version", "?")
	appname := dp.GetProperty("application.name", "?")
	hostname := dp.GetProperty("host.name", "?")

	vlog.Info("Starting [", appname, "] using ", goVer, " on ", hostname)
	vlog.Info("application.profiles.active: ", loading.profile)

	return nil
}
