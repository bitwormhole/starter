package util

import "time"

const (
	theNsPerMs  = 1000000
	theNsPerSec = 1000000000
	theMsPerSec = 1000
)

// CurrentTimestamp 返回当前的时间戳，基于1970-01-01_00:00:00，单位：ms
func CurrentTimestamp() int64 {
	now := time.Now()
	return TimeToInt64(now)
}

// TimeToInt64 返回 time.Time 的 int64(ms) 形式时间戳，基于1970-01-01_00:00:00，单位：ms
func TimeToInt64(t time.Time) int64 {
	sec := t.Unix()
	ns := t.UnixNano() % theNsPerSec
	return (sec * 1000) + (ns / theNsPerMs)
}

// Int64ToTime 是 TimeToInt64 的反函数
func Int64ToTime(ms int64) time.Time {
	sec := ms / theMsPerSec
	ns := (ms % theMsPerSec) * theNsPerMs
	return time.Unix(sec, ns)
}

////////////////////////////////////////////////////////////////////////////////

// Time 用 int64 表示的时间戳，类似java里面的 long System.currentTimeMillis()
type Time int64

// String 转为字符串
func (t Time) String() string {
	t2 := t.Int64()
	t3 := Int64ToTime(t2)
	return t3.String()
}

// Int64 转为整形
func (t Time) Int64() int64 {
	return int64(t)
}

// GetTime 转为 time.Time
func (t Time) GetTime() time.Time {
	n := int64(t)
	return Int64ToTime(n)
}

// NewTimeWithInt64 根据参数创建对应的时间戳
func NewTimeWithInt64(n int64) Time {
	return Time(n)
}

// NewTime 根据参数创建对应的时间戳
func NewTime(t time.Time) Time {
	n := TimeToInt64(t)
	return Time(n)
}

// Now 根据参数创建对应的时间戳
func Now() Time {
	now := time.Now()
	return NewTime(now)
}

////////////////////////////////////////////////////////////////////////////////
