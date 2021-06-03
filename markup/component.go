package markup

// Component  是一个标记，用于配置一个组件
// 请在配置组件的 struct {} 里面包含以下3个字段
//    markup.Component                   // 必须的
//    instance    *foo.Bar               // 必须的
//    context     application.Context    // 可选的
type Component struct {
	// this is a empty struct
}
