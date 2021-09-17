package task

// ThenFn 是 Then 方法的回调函数
type ThenFn func(result interface{})

// CatchFn 是 Catch 方法的回调函数
type CatchFn func(err error)

// ResolveFn 操作成功的处理函数
type ResolveFn func(result interface{})

// RejectFn 是操作失败的处理函数
type RejectFn func(err error)

// FinallyFn 是操作完成（成功|失败）的处理函数
type FinallyFn func()

// PromiseFn 是任务的执行过程
type PromiseFn func(resolve ResolveFn, reject RejectFn)

// Promise 对象用于表示一个异步操作的最终完成 (或失败)及其结果值。
type Promise interface {
	Then(fn ThenFn) Promise
	Catch(fn CatchFn) Promise
	Finally(fn FinallyFn) Promise
}

// NewPromise 新建一个 Promise 对象的实例
func NewPromise(fn PromiseFn) Promise {
	task := &innerPromise{}
	task.start(fn, nil)
	return task
}

// NewPromiseWithExecutor 新建一个 Promise 对象的实例, 并使用指定的 Executor 执行
func NewPromiseWithExecutor(executor Executor, fn PromiseFn) Promise {
	task := &innerPromise{}
	task.start(fn, executor)
	return task
}
