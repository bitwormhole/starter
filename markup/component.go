package markup

////////////////////////////////////////////////////////////////////////////////

// Component  是一个标记，用于配置一个组件
// 请在配置组件的 struct {} 里面包含以下3个字段
//    markup.Component                   // 必须的
//    instance    *foo.Bar               // 必须的
//    context     application.Context    // 可选的
type Component struct {
	// this is a empty struct
}

////////////////////////////////////////////////////////////////////////////////

// Controller  是一个标记，用于配置一个Web控制器组件
type Controller struct {
	Component
}

// RestController  是一个标记，用于配置一个REST控制器组件
type RestController struct {
	Component
	Controller
}

// Repository  是一个标记，用于配置一个存储库组件
type Repository struct {
	Component
}

// DataSource  是一个标记，用于配置一个数据源组件
type DataSource struct {
	Component
}

////////////////////////////////////////////////////////////////////////////////

// IsComponentMark 判断给定的名称是不是一个组件标志
func IsComponentMark(mark string) bool {
	switch mark {
	case "Component":
	case "Controller":
	case "RestController":
	case "Repository":
	case "DataSource":
		break
	default:
		return false
	}
	return true
}

// IsComponentMarkWithPackage 判断给定的名称（以及包名）是不是一个组件标志
func IsComponentMarkWithPackage(pkg string, mark string) bool {
	const regName = "github.com/bitwormhole/starter/markup"
	if pkg != regName {
		return false
	}
	return IsComponentMark(mark)
}
