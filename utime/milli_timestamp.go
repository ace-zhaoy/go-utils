package utime

import "time"

type MilliTimestamp int64

func (m MilliTimestamp) ToTimePtrWithNano(nanoOffset int64, parseZero ...bool) *time.Time {
	if m == 0 && (len(parseZero) == 0 || !parseZero[0]) {
		return nil
	}
	sec := int64(m) / 1e3
	nsec := (int64(m)%1e3)*1e6 + nanoOffset
	tm := time.Unix(sec, nsec)
	return &tm
}

func (m MilliTimestamp) ToTimeWithNano(nanoOffset int64, parseZero ...bool) time.Time {
	tm := m.ToTimePtrWithNano(nanoOffset, parseZero...)
	if tm == nil {
		return time.Time{}
	}
	return *tm
}

func (m MilliTimestamp) ToTimePtr(parseZero ...bool) *time.Time {
	return m.ToTimePtrWithNano(0, parseZero...)
}

func (m MilliTimestamp) ToTime(parseZero ...bool) time.Time {
	return m.ToTimeWithNano(0, parseZero...)
}

func (m MilliTimestamp) ToTimeEndPtr(parseZero ...bool) *time.Time {
	return m.ToTimePtrWithNano(999_999, parseZero...)
}

func (m MilliTimestamp) ToTimeEnd(parseZero ...bool) time.Time {
	return m.ToTimeWithNano(999_999, parseZero...)
}
