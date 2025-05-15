package utime

import "time"

const (
	minNano = 0
	maxNano = 999_999_999
)

type Timestamp int64

func (t Timestamp) ToTimePtrWithNano(nsec int64, parseZero ...bool) *time.Time {
	if t == 0 && (len(parseZero) == 0 || !parseZero[0]) {
		return nil
	}
	tm := time.Unix(int64(t), nsec)
	return &tm
}

func (t Timestamp) ToTimeWithNano(nsec int64, parseZero ...bool) time.Time {
	tm := t.ToTimePtrWithNano(nsec, parseZero...)
	if tm == nil {
		return time.Time{}
	}
	return *tm
}

func (t Timestamp) ToTimePtr(parseZero ...bool) *time.Time {
	return t.ToTimePtrWithNano(minNano, parseZero...)
}

func (t Timestamp) ToTime(parseZero ...bool) time.Time {
	return t.ToTimeWithNano(minNano, parseZero...)
}

func (t Timestamp) ToTimeEndPtr(parseZero ...bool) *time.Time {
	return t.ToTimePtrWithNano(maxNano, parseZero...)
}

func (t Timestamp) ToTimeEnd(parseZero ...bool) time.Time {
	return t.ToTimeWithNano(maxNano, parseZero...)
}
