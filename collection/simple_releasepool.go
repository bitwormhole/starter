package collection

import "github.com/bitwormhole/starter/lang"

type SimpleReleasePool struct {
	list []lang.Disposable
}

func (inst *SimpleReleasePool) Release() []error {
	list := inst.list
	inst.list = nil
	if list == nil {
		return nil
	}
	size := len(list)
	results := make([]error, 0)
	for i := size - 1; i >= 0; i-- {
		item := list[i]
		if item == nil {
			continue
		}
		err := item.Dispose()
		if err != nil {
			results = append(results, err)
		}
	}
	if len(results) == 0 {
		return nil
	}
	return results
}

func (inst *SimpleReleasePool) Push(target lang.Disposable) {
	if target == nil {
		return
	}
	list := inst.list
	if list == nil {
		list = make([]lang.Disposable, 0)
	}
	inst.list = append(list, target)
}

func CreateReleasePool() ReleasePool {
	inst := &SimpleReleasePool{}
	return inst
}
