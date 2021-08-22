package util

import "time"

// CurrentTimestamp 返回当前的时间戳，基于1970-01-01_00:00:00，单位：ms
func CurrentTimestamp() int64 {
	now := time.Now()
	sec := now.Unix()
	// now.UnixNano()
	return sec * 1000
}
