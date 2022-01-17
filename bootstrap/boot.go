package bootstrap

import (
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// Boot 主启动器
type Boot struct {
	markup.Component `id:"main-looper"`

	Lives      []lang.Object `inject:".life"`
	Concurrent bool          `inject:"${application.loopers.concurrent}"`

	proxies []*lifeProxy
}

func (inst *Boot) _Impl() application.MainLooper {
	return inst
}

// RunMain ...
func (inst *Boot) RunMain() error {
	err := inst.load()
	if err != nil {
		return err
	}
	err = inst.start()
	if err != nil {
		vlog.Error(err)
	} else {
		err = inst.loop()
		if err != nil {
			vlog.Error(err)
		}
	}
	return inst.stop()
}

func (inst *Boot) load() error {
	src := inst.Lives
	dst := make([]*lifeProxy, 0)
	for _, item := range src {
		proxy := &lifeProxy{parent: inst}
		cnt := proxy.load(item)
		if cnt > 0 {
			dst = append(dst, proxy)
		}
	}
	inst.proxies = dst
	return nil
}

func (inst *Boot) start() error {
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

func (inst *Boot) listLoopers() []application.Looper {
	src := inst.proxies
	dst := make([]application.Looper, 0)
	for _, item := range src {
		looper := item.mLooper.Looper
		if looper != nil {
			dst = append(dst, looper)
		}
	}
	return dst
}

func (inst *Boot) loop() error {
	if inst.Concurrent {
		return inst.loopAsConcurrent()
	}
	return inst.loopAsSerial()
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
		runner.start(item)
	}
	return runner.waitForAllDone()
}

func (inst *Boot) stop() error {
	all := inst.proxies
	i := len(all)
	for i--; i >= 0; i-- {
		item := all[i]
		err := item.stop()
		if err != nil {
			vlog.Error(err)
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type lifeProxy struct {
	parent   *Boot
	mStarter application.StarterRegistration
	mLooper  application.LooperRegistration
	mStopper application.StopperRegistration
	started  bool
}

func (inst *lifeProxy) load(o lang.Object) int {

	count := 0

	o1, ok := o.(application.StarterRegistry)
	if ok {
		inst.mStarter = *o1.GetStarterRegistration()
		count++
	}

	o2, ok := o.(application.LooperRegistry)
	if ok {
		inst.mLooper = *o2.GetLooperRegistration()
		count++
	}

	o3, ok := o.(application.StopperRegistry)
	if ok {
		inst.mStopper = *o3.GetStopperRegistration()
		count++
	}

	return count
}

func (inst *lifeProxy) start() error {
	target := inst.mStarter.Starter
	if target != nil {
		err := target.Start()
		if err != nil {
			return err
		}
	}
	inst.started = true
	return nil
}

func (inst *lifeProxy) loop() error {

	if !inst.started {
		return nil
	}

	target := inst.mLooper.Looper
	if target == nil {
		return nil
	}
	return target.Loop()
}

func (inst *lifeProxy) stop() error {

	if !inst.started {
		return nil
	}

	target := inst.mStopper.Stopper
	if target == nil {
		return nil
	}
	return target.Stop()
}

////////////////////////////////////////////////////////////////////////////////

type concurrentLooperRunner struct {
	countBegin int
	countEnd   int
}

func (inst *concurrentLooperRunner) waitForAllDone() error {
	for {
		if inst.countEnd < inst.countBegin {
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}
	return nil
}

func (inst *concurrentLooperRunner) start(looper application.Looper) {
	if looper == nil {
		return
	}
	inst.countBegin++
	go func() {
		inst.run(looper)
	}()
}

func (inst *concurrentLooperRunner) run(looper application.Looper) {
	defer func() {
		e := recover()
		if e != nil {
			vlog.Error(e)
		}
	}()
	defer func() {
		inst.incEnd()
	}()
	err := looper.Loop()
	if err != nil {
		vlog.Error(err)
	}
}

func (inst *concurrentLooperRunner) incEnd() {
	// TODO:  mutex ...
	inst.countEnd++
}
