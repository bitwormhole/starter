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
