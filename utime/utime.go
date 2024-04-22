package utime

import "time"

func Unix(sec int64) time.Time {
	if sec == 0 {
		return time.Time{}
	}
	return time.Unix(sec, 0)
}

func UnixMilli(msec int64) time.Time {
	if msec == 0 {
		return time.Time{}
	}
	return time.UnixMilli(msec)
}
