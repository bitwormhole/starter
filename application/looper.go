package application

////////////////////////////////////////////////////////////////////////////////

// Looper 【inject:"#main-looper"】
type Looper interface {
	Loop() error
}

////////////////////////////////////////////////////////////////////////////////

// MainLooper 【inject:"#main-looper"】
type MainLooper interface {
	Killer
	RunMain() error
}

////////////////////////////////////////////////////////////////////////////////

func runMainLooper(c Context) error {
	const selector = "#main-looper"
	o1, err := c.GetComponent(selector)
	if err != nil {
		return err
	}
	o2 := o1.(MainLooper)
	return o2.RunMain()
}

// GetMainLooper 取主循环对象
func GetMainLooper(c Context) MainLooper {
	const selector = "#main-looper"
	o1, err := c.GetComponent(selector)
	if err != nil {
		panic(err)
	}
	o2 := o1.(MainLooper)
	return o2
}

////////////////////////////////////////////////////////////////////////////////
// EOF
