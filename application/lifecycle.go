package application

import "sort"

// app 生命周期： Init -> Start -> Loop ->  Stop -> Destroy!

// 定义app生命周期组件的类名称
const (
	LifeClassName = "life" // 	统一使用'class:"life"'注入
)

// OnLifeFunc 定义生命周期处理函数
type OnLifeFunc func() error

// LifeRegistry 表示 Life 的注册者
type LifeRegistry interface {
	GetLifeRegistration() *LifeRegistration
}

// LifeRegistration 表示 Life 的注册项
type LifeRegistration struct {
	Priority int // 数值越大，优先级越高，优先执行OnInit & OnStart

	OnInit    OnLifeFunc
	OnStart   OnLifeFunc
	Looper    Looper
	OnStop    OnLifeFunc
	OnDestroy OnLifeFunc
}

// Life 是 LifeRegistration 的别名
type Life LifeRegistration

// Killer 接口用于通知应用程序关闭,  【inject:".killer"】
type Killer interface {
	Shutdown() error
}

////////////////////////////////////////////////////////////////////////////////

// LifeRegistrationSorter 是 LifeRegistration 的排序器
type LifeRegistrationSorter struct {
	List []*LifeRegistration
}

func (inst *LifeRegistrationSorter) _Impl() sort.Interface {
	return inst
}

func (inst *LifeRegistrationSorter) Len() int {
	list := inst.List
	if list == nil {
		return 0
	}
	return len(list)
}

func (inst *LifeRegistrationSorter) Less(a, b int) bool {
	list := inst.List
	return list[a].Priority > list[b].Priority
}

func (inst *LifeRegistrationSorter) Swap(a, b int) {
	list := inst.List
	aa := list[a]
	bb := list[b]
	list[a] = bb
	list[b] = aa
}

////////////////////////////////////////////////////////////////////////////////
