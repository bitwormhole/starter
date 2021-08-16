package tests

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/starter/io/fs"
)

type DefaultRunner struct {
	indexForCase int
	baseTempDir  fs.Path
}

func (inst *DefaultRunner) _Impl() TestRunner {
	return inst
}

func (inst *DefaultRunner) Run(ctx TestContext) error {
	if inst.baseTempDir == nil {
		inst.baseTempDir = ctx.TempDir()
	}
	all := ctx.CaseManager().All()
	for _, info := range all {
		inst.tryRunInfo(info, ctx)
	}
	return inst.logResults(ctx, all)
}

func (inst *DefaultRunner) logResults(ctx TestContext, list []*CaseInfo) error {
	logger := ctx.Logger()
	var err error = nil
	for _, info := range list {
		if info.Error == nil {
			continue
		}
		err = info.Error
		// log error
		logger.Error("[TestCase id:", info.ID, " class:", info.Class, " error:", err.Error(), "]")
	}
	return err
}

func (inst *DefaultRunner) now() int64 {
	return time.Now().Unix() * 1000
}

func (inst *DefaultRunner) tryRunInfo(info *CaseInfo, ctx TestContext) {
	if info.Done {
		return
	} else if info.TimeBegin > 0 {
		return
	}

	t1 := inst.now()
	err := inst.inTryRunInfo(info, ctx)
	if err != nil {
		info.Error = err
	}

	info.TimeBegin = t1
	info.TimeEnd = inst.now()
	info.Done = true
}

func (inst *DefaultRunner) inTryRunInfo(info *CaseInfo, ctx TestContext) error {

	defer func() {
		p := recover()
		if p == nil {
			return
		}
		msg := fmt.Sprint("panic:", p)
		info.Error = errors.New(msg)
	}()

	ctx2 := &DefaultTestContext{}
	ctx2.InitWith(ctx)
	ctx2.MyTempDir = inst.computeTempDir(info)

	return info.Case.OnTest(ctx2)
}

func (inst *DefaultRunner) computeTempDir(info *CaseInfo) fs.Path {

	inst.indexForCase++
	index := inst.indexForCase
	now := inst.now()
	id := info.ID

	builder := strings.Builder{}
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString("-")
	builder.WriteString(strconv.FormatInt(now, 10))
	builder.WriteString("-case")
	builder.WriteString(id)
	builder.WriteString(".t")

	return inst.baseTempDir.GetChild(builder.String())
}
