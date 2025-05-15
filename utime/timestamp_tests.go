package utime

import (
	"testing"
)

func TestTimestamp_ToTimePtrWithNano(t *testing.T) {
	tests := []struct {
		name       string
		timestamp  Timestamp
		nsec       int64
		parseZero  []bool
		wantNotNil bool
	}{
		{"NonZeroTimestamp", 1234567890, 123456789, nil, true},
		{"ZeroTimestampNoParse", 0, 0, nil, false},
		{"ZeroTimestampParse", 0, 0, []bool{true}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.timestamp.ToTimePtrWithNano(tt.nsec, tt.parseZero...)
			if (result != nil) != tt.wantNotNil {
				t.Errorf("ToTimePtrWithNano() = %v, wantNotNil %v", result, tt.wantNotNil)
			}
		})
	}
}

func TestTimestamp_ToTimeWithNano(t *testing.T) {
	tests := []struct {
		name      string
		timestamp Timestamp
		nsec      int64
		parseZero []bool
		wantZero  bool
	}{
		{"NonZeroTimestamp", 1234567890, 123456789, nil, false},
		{"ZeroTimestampNoParse", 0, 0, nil, true},
		{"ZeroTimestampParse", 0, 0, []bool{true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.timestamp.ToTimeWithNano(tt.nsec, tt.parseZero...)
			if (result.IsZero()) != tt.wantZero {
				t.Errorf("ToTimeWithNano() = %v, wantZero %v", result, tt.wantZero)
			}
		})
	}
}

func TestTimestamp_ToTimePtr(t *testing.T) {
	tests := []struct {
		name       string
		timestamp  Timestamp
		parseZero  []bool
		wantNotNil bool
	}{
		{"NonZeroTimestamp", 1234567890, nil, true},
		{"ZeroTimestampNoParse", 0, nil, false},
		{"ZeroTimestampParse", 0, []bool{true}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.timestamp.ToTimePtr(tt.parseZero...)
			if (result != nil) != tt.wantNotNil {
				t.Errorf("ToTimePtr() = %v, wantNotNil %v", result, tt.wantNotNil)
			}
		})
	}
}

func TestTimestamp_ToTime(t *testing.T) {
	tests := []struct {
		name      string
		timestamp Timestamp
		parseZero []bool
		wantZero  bool
	}{
		{"NonZeroTimestamp", 1234567890, nil, false},
		{"ZeroTimestampNoParse", 0, nil, true},
		{"ZeroTimestampParse", 0, []bool{true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.timestamp.ToTime(tt.parseZero...)
			if (result.IsZero()) != tt.wantZero {
				t.Errorf("ToTime() = %v, wantZero %v", result, tt.wantZero)
			}
		})
	}
}

func TestTimestamp_ToTimeEndPtr(t *testing.T) {
	tests := []struct {
		name       string
		timestamp  Timestamp
		parseZero  []bool
		wantNotNil bool
	}{
		{"NonZeroTimestamp", 1234567890, nil, true},
		{"ZeroTimestampNoParse", 0, nil, false},
		{"ZeroTimestampParse", 0, []bool{true}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.timestamp.ToTimeEndPtr(tt.parseZero...)
			if (result != nil) != tt.wantNotNil {
				t.Errorf("ToTimeEndPtr() = %v, wantNotNil %v", result, tt.wantNotNil)
			}
		})
	}
}

func TestTimestamp_ToTimeEnd(t *testing.T) {
	tests := []struct {
		name      string
		timestamp Timestamp
		parseZero []bool
		wantZero  bool
	}{
		{"NonZeroTimestamp", 1234567890, nil, false},
		{"ZeroTimestampNoParse", 0, nil, true},
		{"ZeroTimestampParse", 0, []bool{true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.timestamp.ToTimeEnd(tt.parseZero...)
			if (result.IsZero()) != tt.wantZero {
				t.Errorf("ToTimeEnd() = %v, wantZero %v", result, tt.wantZero)
			}
		})
	}
}
