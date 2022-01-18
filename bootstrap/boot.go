package bootstrap

import (
	"sort"
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// Boot 主启动器
type Boot struct {
	markup.Component `id:"main-looper"`

	Lives      []application.LifeRegistry `inject:".life"`
	Concurrent bool                       `inject:"${application.loopers.concurrent}"`

	proxies []*lifeProxy
}

func (inst *Boot) _Impl() application.MainLooper {
	return inst
}

func (inst *Boot) logError(e error) {
	if e == nil {
		return
	}
	vlog.Error(e)
}

// RunMain ...
func (inst *Boot) RunMain() error {
	err := inst.prepare()
	if err != nil {
		return err
	}
	return inst.run2()
}

func (inst *Boot) run2() error {
	err := inst.doInit()
	defer func() {
		e2 := inst.doDestroy()
		inst.logError(e2)
	}()
	if err != nil {
		return err
	}
	return inst.run3()
}

func (inst *Boot) run3() error {
	err := inst.doStart()
	defer func() {
		e2 := inst.doStop()
		inst.logError(e2)
	}()
	if err != nil {
		return err
	}
	return inst.run4()
}

func (inst *Boot) run4() error {
	defer func() {
		e := recover()
		if e != nil {
			vlog.Error(e)
		}
	}()
	return inst.doLoop()
}

func (inst *Boot) prepare() error {

	src := inst.Lives
	dst := make([]*lifeProxy, 0)
	mid := make([]*application.LifeRegistration, 0)

	for _, item := range src {
		lr := item.GetLifeRegistration()
		if lr == nil {
			continue
		}
		mid = append(mid, lr)
	}

	sort.Sort(&application.LifeRegistrationSorter{List: mid})

	for _, item := range mid {
		proxy := &lifeProxy{parent: inst}
		proxy.prepare(item)
		dst = append(dst, proxy)
	}

	inst.proxies = dst
	return nil
}

func (inst *Boot) doInit() error {
	all := inst.proxies
	for _, item := range all {
		err := item.init()
		if err != nil {
			return err
		}
		item.initialled = true
	}
	return nil
}

func (inst *Boot) doStart() error {
	all := inst.proxies
	for _, item := range all {
		err := item.start()
		if err != nil {
			return err
		}
		item.started = true
	}
	return nil
}

func (inst *Boot) doLoop() error {
	if inst.Concurrent {
		return inst.loopAsConcurrent()
	}
	return inst.loopAsSerial()
}

func (inst *Boot) doStop() error {
	all := inst.proxies
	i := len(all) - 1
	for ; i >= 0; i-- {
		item := all[i]
		err := item.stop()
		if err != nil {
			vlog.Error(err)
		}
	}
	return nil
}

func (inst *Boot) doDestroy() error {
	all := inst.proxies
	i := len(all) - 1
	for ; i >= 0; i-- {
		item := all[i]
		err := item.destroy()
		if err != nil {
			vlog.Error(err)
		}
	}
	return nil
}

func (inst *Boot) loopAsSerial() error {
	all := inst.listLoopers()
	for _, item := range all {
		err := item.Loop()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *Boot) loopAsConcurrent() error {
	all := inst.listLoopers()
	runner := concurrentLooperRunner{}
	for _, item := range all {
		runner.startLooper(item)
	}
	return runner.waitForAllDone()
}

func (inst *Boot) listLoopers() []application.Looper {
	src := inst.proxies
	dst := make([]application.Looper, 0)
	for _, item := range src {
		looper := item.Looper
		if looper != nil {
			dst = append(dst, looper)
		}
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type lifeProxy struct {
	application.LifeRegistration

	parent *Boot

	initialled bool
	started    bool
}

func (inst *lifeProxy) prepare(o *application.LifeRegistration) {
	if o == nil {
		return
	}
	inst.LifeRegistration = *o
}

func (inst *lifeProxy) init() error {
	fn := inst.OnInit
	if fn == nil {
		return nil
	}
	return fn()
}

func (inst *lifeProxy) start() error {
	fn := inst.OnStart
	if fn == nil {
		return nil
	}
	return fn()
}

func (inst *lifeProxy) loop() error {
	if !inst.started {
		return nil
	}
	looper := inst.Looper
	if looper == nil {
		return nil
	}
	return looper.Loop()
}

func (inst *lifeProxy) stop() error {
	if !inst.started {
		return nil
	}
	fn := inst.OnStop
	if fn == nil {
		return nil
	}
	return fn()
}

func (inst *lifeProxy) destroy() error {
	if !inst.initialled {
		return nil
	}
	fn := inst.OnDestroy
	if fn == nil {
		return nil
	}
	return fn()
}

////////////////////////////////////////////////////////////////////////////////

type concurrentLooperRunner struct {
	all []*concurrentLooperHolder
}

func (inst *concurrentLooperRunner) startLooper(l application.Looper) {
	if l == nil {
		return
	}
	h := &concurrentLooperHolder{}
	h.parent = inst
	h.looper = l
	inst.all = append(inst.all, h)
	h.start()
}

func (inst *concurrentLooperRunner) isAllDone() bool {
	all := inst.all
	if all == nil {
		return true
	}
	for _, item := range all {
		if !item.stopped {
			return false
		}
	}
	return true
}

func (inst *concurrentLooperRunner) waitForAllDone() error {
	for {
		if !inst.isAllDone() {
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}
	return nil
}

////////////////////

type concurrentLooperHolder struct {
	parent  *concurrentLooperRunner
	started bool
	stopped bool
	looper  application.Looper
}

func (inst *concurrentLooperHolder) Loop() error {
	return nil // NOP
}

func (inst *concurrentLooperHolder) start() {
	looper := inst.looper
	if looper == nil {
		looper = inst // set as NOP
	}
	go func() {
		inst.run(looper)
	}()
}

func (inst *concurrentLooperHolder) run(looper application.Looper) {
	defer func() {
		e := recover()
		if e != nil {
			vlog.Error(e)
		}
	}()
	inst.started = true
	defer func() {
		inst.stopped = true
	}()
	err := looper.Loop()
	if err != nil {
		vlog.Error(err)
	}
}
