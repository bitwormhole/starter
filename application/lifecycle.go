package application

import "sort"

// 定义app生命周期组件的类名称
const (
	// StarterClassName = "starter"
	// LooperClassName  = "looper"
	// StopperClassName = "stopper"
	LifeClassName = "life" // 	统一使用“life”注入
)

// app 生命周期： Init -> Start -> Loop ->  Stop -> Destroy!

////////////////////////////////////////////////////////////////////////////////
// Init - 由 markup.Component 自定义

////////////////////////////////////////////////////////////////////////////////
// Starter

// Starter 是 app 的启动器
type Starter interface {
	Start() error
}

// StarterRegistry 表示 Looper 的注册者
type StarterRegistry interface {
	GetStarterRegistration() *StarterRegistration
}

// StarterRegistration 表示 Starter 的注册项
type StarterRegistration struct {
	Priority int
	Starter  Starter
}

// StarterRegistrationSorter 是 StarterRegistration 的排序器
type StarterRegistrationSorter struct {
	List []*StarterRegistration
}

func (inst *StarterRegistrationSorter) _Impl() sort.Interface {
	return inst
}

func (inst *StarterRegistrationSorter) Len() int {
	list := inst.List
	if list == nil {
		return 0
	}
	return len(list)
}

func (inst *StarterRegistrationSorter) Less(a, b int) bool {
	list := inst.List
	return list[a].Priority < list[b].Priority
}

func (inst *StarterRegistrationSorter) Swap(a, b int) {
	list := inst.List
	aa := list[a]
	bb := list[b]
	list[a] = bb
	list[b] = aa
}

////////////////////////////////////////////////////////////////////////////////
// Looper

// Looper 是 app 的循环器
type Looper interface {
	Loop() error
}

// LooperRegistry 表示 Looper 的注册者
type LooperRegistry interface {
	GetLooperRegistration() *LooperRegistration
}

// LooperRegistration 表示 Looper 的注册项
type LooperRegistration struct {
	Priority int
	Looper   Looper
}

// LooperRegistrationSorter 是 LooperRegistration 的排序器
type LooperRegistrationSorter struct {
	List []*LooperRegistration
}

func (inst *LooperRegistrationSorter) _Impl() sort.Interface {
	return inst
}

func (inst *LooperRegistrationSorter) Len() int {
	list := inst.List
	if list == nil {
		return 0
	}
	return len(list)
}

func (inst *LooperRegistrationSorter) Less(a, b int) bool {
	list := inst.List
	return list[a].Priority < list[b].Priority
}

func (inst *LooperRegistrationSorter) Swap(a, b int) {
	list := inst.List
	aa := list[a]
	bb := list[b]
	list[a] = bb
	list[b] = aa
}

////////////////////////////////////////////////////////////////////////////////
// Stopper

// Stopper 是 app 的制动器
type Stopper interface {
	Stop() error
}

// StopperRegistry 表示 Looper 的注册者
type StopperRegistry interface {
	GetStopperRegistration() *StopperRegistration
}

// StopperRegistration 表示 Stopper 的注册项
type StopperRegistration struct {
	Priority int
	Stopper  Stopper
}

// StopperRegistrationSorter 是 StopperRegistration 的排序器
type StopperRegistrationSorter struct {
	List []*StopperRegistration
}

func (inst *StopperRegistrationSorter) _Impl() sort.Interface {
	return inst
}

func (inst *StopperRegistrationSorter) Len() int {
	list := inst.List
	if list == nil {
		return 0
	}
	return len(list)
}

func (inst *StopperRegistrationSorter) Less(a, b int) bool {
	list := inst.List
	return list[a].Priority < list[b].Priority
}

func (inst *StopperRegistrationSorter) Swap(a, b int) {
	list := inst.List
	aa := list[a]
	bb := list[b]
	list[a] = bb
	list[b] = aa
}

////////////////////////////////////////////////////////////////////////////////
// Destroy  - 由 markup.Component 自定义

////////////////////////////////////////////////////////////////////////////////

// Lifecycle 是完整的app生命周期
type Lifecycle interface {
	Starter
	Looper
	Stopper
}
