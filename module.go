package starter

import (
	"github.com/bitwormhole/starter/application"
	etc "github.com/bitwormhole/starter/etc/starter"
)

// Module 导出【starter】模块
func Module() application.Module {
	return etc.ExportModule()
}
