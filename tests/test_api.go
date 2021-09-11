package tests

import (
	"embed"
	"testing"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
)

// TestingInitializer 是对 application.Initializer 的扩展，添加了几个用于测试的功能
type TestingInitializer interface {
	application.Initializer

	T() *testing.T

	UseResourcesFS(efs *embed.FS, path string) TestingInitializer
	PrepareTestingDataFromResource(name string) TestingInitializer
	LoadPropertisFromGitConfig(required bool) TestingInitializer

	RunTest() TestingRuntime
}

// TestingRuntime 是对 application.Runtime 的扩展，添加了几个用于测试的功能
type TestingRuntime interface {
	application.Runtime

	T() *testing.T
	TestingDataDir() fs.Path
}
