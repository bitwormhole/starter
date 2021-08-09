package platforms

// Platform 定义一个跨平台的统一接口
type Platform interface {
	OS() OS
}

// OS 接口用于获取操作系统信息
type OS interface {
	Name() string
	Version() string
}

// GetPlatform 函数用于获取 Platform 接口的实例
func GetPlatform() Platform {
	// todo ...
	return nil
}
