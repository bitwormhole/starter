package application

// Initializer 是应用程序的启动器
type Initializer interface {

	// EmbedResources(fs *embed.FS, path string) Initializer
	// MountResources(res collection.Resources) Initializer

	SetAttribute(name string, value interface{}) Initializer
	Use(module Module) Initializer
	Run()
}
