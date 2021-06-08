package runtime

import (
	"errors"

	"github.com/bitwormhole/starter/lang"
)

type PoolPairBinding struct {
	pool1 lang.ReleasePool
	pool2 lang.ReleasePool
}

func (inst *PoolPairBinding) Init(p1 lang.ReleasePool, p2 lang.ReleasePool) error {

	if p1 == nil {
		return errors.New("pool_1==nil")
	}
	if p2 == nil {
		return errors.New("pool_2==nil")
	}

	inst.pool1 = p1
	inst.pool2 = p2
	p1.Push(inst)
	p2.Push(inst)
	return nil
}

func (inst *PoolPairBinding) Dispose() error {

	p1 := inst.pool1
	p2 := inst.pool2
	var err1 error
	var err2 error

	inst.pool1 = nil
	inst.pool2 = nil

	if p1 != nil {
		err1 = p1.Release()
	}
	if p2 != nil {
		err2 = p2.Release()
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}
