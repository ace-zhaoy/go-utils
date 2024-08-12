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

func StartOfMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()
	return time.Date(year, month, day, hour, minute, 0, 0, t.Location())
}

func StartOfHour(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, _, _ := t.Clock()
	return time.Date(year, month, day, hour, 0, 0, 0, t.Location())
}

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func StartOfWeek(t time.Time) time.Time {
	year, month, day := t.Date()
	wd := int(t.Weekday()) - 1
	if wd < 0 {
		wd = 6
	}

	return time.Date(year, month, day-wd, 0, 0, 0, 0, t.Location())
}

func StartOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}
