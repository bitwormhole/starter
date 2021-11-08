package task

import "github.com/bitwormhole/starter/vlog"

type fireEvent int

const (
	fireEventThen    fireEvent = 1
	fireEventCatch   fireEvent = 2
	fireEventFinally fireEvent = 3
)

////////////////////////////////////////////////////////////////////////////////

type callbackChainNode struct {
	fnThen    ThenFn
	fnCatch   CatchFn
	fnFinally FinallyFn

	wantEvent fireEvent
	fired     bool
	promise   *innerPromise
}

func (inst *callbackChainNode) fire(event fireEvent) {

	if event != inst.wantEvent {
		return
	}

	if inst.fired {
		return
	}
	inst.fired = true

	switch event {
	case fireEventThen:
		inst.tryFireThen()
		break
	case fireEventCatch:
		inst.tryFireCatch()
		break
	case fireEventFinally:
		inst.tryFireFinally()
		break
	}
}

func (inst *callbackChainNode) tryFireThen() {
	callback := inst.fnThen
	if callback == nil {
		return
	}
	callback(inst.promise.result)
}

func (inst *callbackChainNode) tryFireCatch() {

	defer func() {
		o := recover()
		if o != nil {
			vlog.Warn("recover: ", o)
		}
	}()

	callback := inst.fnCatch
	if callback == nil {
		return
	}
	callback(inst.promise.err)
}

func (inst *callbackChainNode) tryFireFinally() {

	defer func() {
		o := recover()
		if o != nil {
			vlog.Warn("recover: ", o)
		}
	}()

	callback := inst.fnFinally
	if callback == nil {
		return
	}
	callback()
}

////////////////////////////////////////////////////////////////////////////////

type innerPromise struct {
	callbackChain []*callbackChainNode
	result        interface{}
	err           error
	task          PromiseFn
	done          bool
}

func (inst *innerPromise) _Impl() Promise {
	return inst
}

func (inst *innerPromise) append(node *callbackChainNode) {
	if node == nil {
		return
	}
	node.promise = inst
	if inst.done {
		inst.fireNodeAfterDone(node)
		return
	}
	inst.callbackChain = append(inst.callbackChain, node)
	if inst.done {
		inst.fireAll(node.wantEvent)
	}
}

func (inst *innerPromise) fireNodeAfterDone(node *callbackChainNode) {

	err := inst.err
	event := node.wantEvent

	switch event {
	case fireEventThen:
		if err == nil {
			node.fire(event)
		}
		break
	case fireEventCatch:
		if err != nil {
			node.fire(event)
		}
		break
	case fireEventFinally:
		node.fire(event)
		break
	}
}

func (inst *innerPromise) fireAll(event fireEvent) {
	all := inst.callbackChain
	if all == nil {
		return
	}
	for _, node := range all {
		node.fire(event)
	}
}

func (inst *innerPromise) Then(fn ThenFn) Promise {
	node := &callbackChainNode{
		fnThen:    fn,
		wantEvent: fireEventThen,
	}
	inst.append(node)
	return inst
}

func (inst *innerPromise) Catch(fn CatchFn) Promise {
	node := &callbackChainNode{
		fnCatch:   fn,
		wantEvent: fireEventCatch,
	}
	inst.append(node)
	return inst
}

func (inst *innerPromise) Finally(fn FinallyFn) Promise {
	node := &callbackChainNode{
		fnFinally: fn,
		wantEvent: fireEventFinally,
	}
	inst.append(node)
	return inst
}

func (inst *innerPromise) start(fn PromiseFn, executor Executor) {
	if fn == nil {
		return
	}
	inst.task = fn
	if executor == nil {
		executor = DefaultExecutor()
	}
	executor.Execute(inst)
}

func (inst *innerPromise) Run() {

	defer func() {
		o := recover()
		if o != nil {
			vlog.Warn("recover: ", o)
		}
	}()

	defer func() {
		inst.handleFinally()
	}()

	var resolve ResolveFn = func(result interface{}) {
		inst.handleResult(result)
	}

	var reject RejectFn = func(err error) {
		inst.handleError(err)
	}

	inst.task(resolve, reject)
}

func (inst *innerPromise) handleError(err error) {
	inst.err = err
	inst.done = true
	inst.fireAll(fireEventCatch)
}

func (inst *innerPromise) handleResult(result interface{}) {
	inst.result = result
	inst.done = true
	inst.fireAll(fireEventThen)
}

func (inst *innerPromise) handleFinally() {
	inst.done = true
	inst.fireAll(fireEventFinally)
}
