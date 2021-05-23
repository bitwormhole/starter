package lang

type SimpleReleasePool struct {
	list []Disposable
}

func (inst *SimpleReleasePool) Release() error {
	list := inst.list
	inst.list = nil

	if list == nil {
		return nil
	}

	results := NewErrorCollector()
	size := len(list)

	for i := size - 1; i >= 0; i-- {
		item := list[i]
		if item == nil {
			continue
		}
		err := item.Dispose()
		if err != nil {
			results.Append(err)
		}
	}

	return results.Result()
}

func (inst *SimpleReleasePool) Push(target Disposable) {
	if target == nil {
		return
	}
	list := inst.list
	if list == nil {
		list = make([]Disposable, 0)
	}
	inst.list = append(list, target)
}

func CreateReleasePool() ReleasePool {
	inst := &SimpleReleasePool{}
	return inst
}
