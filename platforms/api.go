package platforms

// Platform 定义一个跨平台的统一接口
type Platform interface {

	// 返回系统对象
	GetOS() OS

	// 返回系统名称
	OS() string

	// 返回架构名称
	Arch() string
}

// OS 接口用于获取操作系统信息
type OS interface {
	Name() string
	Version() string
}

////////////////////////////////////////////////////////////////////////////////

var theCurrentPlatform Platform

// Current 函数用于获取 Platform 接口的实例
func Current() Platform {
	p := theCurrentPlatform
	if p == nil {
		p = initCurrent()
		theCurrentPlatform = p
	}
	return p
}
