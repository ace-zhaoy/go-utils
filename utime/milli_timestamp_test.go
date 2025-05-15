package utime

import (
	"testing"
	"time"
)

func TestMilliTimestamp_ToTimePtrWithNano(t *testing.T) {
	var zero MilliTimestamp = 0
	var ts MilliTimestamp = 1680000000123 // 2023-03-28 09:20:00.123 UTC

	// 零值，parseZero未传，返回nil
	if zero.ToTimePtrWithNano(0) != nil {
		t.Error("expected nil for zero timestamp without parseZero")
	}
	// 零值，parseZero=false，返回nil
	if zero.ToTimePtrWithNano(0, false) != nil {
		t.Error("expected nil for zero timestamp with parseZero=false")
	}
	// 零值，parseZero=true，返回非nil
	if zero.ToTimePtrWithNano(0, true) == nil {
		t.Error("expected non-nil for zero timestamp with parseZero=true")
	}
	// 正常值，nanoOffset=0
	tm := ts.ToTimePtrWithNano(0)
	want := time.Unix(1680000000, 123*1e6)
	if tm == nil || !tm.Equal(want) {
		t.Errorf("expected %v, got %v", want, tm)
	}
	// 正常值，nanoOffset=500
	tm = ts.ToTimePtrWithNano(500)
	want = time.Unix(1680000000, 123*1e6+500)
	if tm == nil || !tm.Equal(want) {
		t.Errorf("expected %v, got %v", want, tm)
	}
}

func TestMilliTimestamp_ToTimeWithNano(t *testing.T) {
	var zero MilliTimestamp = 0
	var ts MilliTimestamp = 1680000000123

	// 零值，parseZero未传，返回time.Time{}
	got := zero.ToTimeWithNano(0)
	if !got.IsZero() {
		t.Error("expected zero time for zero timestamp")
	}
	// 零值，parseZero=true，返回非零
	got = zero.ToTimeWithNano(0, true)
	if got.IsZero() {
		t.Error("expected non-zero time for zero timestamp with parseZero=true")
	}
	// 正常值
	got = ts.ToTimeWithNano(0)
	want := time.Unix(1680000000, 123*1e6)
	if !got.Equal(want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestMilliTimestamp_ToTimePtr(t *testing.T) {
	var zero MilliTimestamp = 0
	var ts MilliTimestamp = 1680000000123

	// 零值
	if zero.ToTimePtr() != nil {
		t.Error("expected nil for zero timestamp")
	}
	// 零值，parseZero=true
	if zero.ToTimePtr(true) == nil {
		t.Error("expected non-nil for zero timestamp with parseZero=true")
	}
	// 正常值
	tm := ts.ToTimePtr()
	want := time.Unix(1680000000, 123*1e6)
	if tm == nil || !tm.Equal(want) {
		t.Errorf("expected %v, got %v", want, tm)
	}
}

func TestMilliTimestamp_ToTime(t *testing.T) {
	var zero MilliTimestamp = 0
	var ts MilliTimestamp = 1680000000123

	// 零值
	got := zero.ToTime()
	if !got.IsZero() {
		t.Error("expected zero time for zero timestamp")
	}
	// 零值，parseZero=true
	got = zero.ToTime(true)
	if got.IsZero() {
		t.Error("expected non-zero time for zero timestamp with parseZero=true")
	}
	// 正常值
	want := time.Unix(1680000000, 123*1e6)
	got = ts.ToTime()
	if !got.Equal(want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestMilliTimestamp_ToTimeEndPtr(t *testing.T) {
	var zero MilliTimestamp = 0
	var ts MilliTimestamp = 1680000000123

	// 零值
	if zero.ToTimeEndPtr() != nil {
		t.Error("expected nil for zero timestamp")
	}
	// 零值，parseZero=true
	if zero.ToTimeEndPtr(true) == nil {
		t.Error("expected non-nil for zero timestamp with parseZero=true")
	}
	// 正常值
	tm := ts.ToTimeEndPtr()
	want := time.Unix(1680000000, 123*1e6+999_999)
	if tm == nil || !tm.Equal(want) {
		t.Errorf("expected %v, got %v", want, tm)
	}
}

func TestMilliTimestamp_ToTimeEnd(t *testing.T) {
	var zero MilliTimestamp = 0
	var ts MilliTimestamp = 1680000000123

	// 零值
	got := zero.ToTimeEnd()
	if !got.IsZero() {
		t.Error("expected zero time for zero timestamp")
	}
	// 零值，parseZero=true
	got = zero.ToTimeEnd(true)
	if got.IsZero() {
		t.Error("expected non-zero time for zero timestamp with parseZero=true")
	}
	// 正常值
	want := time.Unix(1680000000, 123*1e6+999_999)
	got = ts.ToTimeEnd()
	if !got.Equal(want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}
