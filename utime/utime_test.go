package utime

import (
	"reflect"
	"testing"
	"time"
)

func TestUnix(t *testing.T) {
	type args struct {
		sec int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test 1",
			args: args{sec: 1},
			want: time.Unix(1, 0),
		},
		{
			name: "test 0",
			args: args{sec: 0},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unix(tt.args.sec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnixMilli(t *testing.T) {
	type args struct {
		msec int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test 1",
			args: args{msec: 1},
			want: time.UnixMilli(1),
		},
		{
			name: "test 0",
			args: args{msec: 0},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnixMilli(tt.args.msec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnixMilli() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfMinute(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Regular Time",
			input:    time.Date(2024, 8, 12, 15, 45, 30, 0, time.UTC),
			expected: time.Date(2024, 8, 12, 15, 45, 0, 0, time.UTC),
		},
		{
			name:     "Edge Time",
			input:    time.Date(2024, 8, 12, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2024, 8, 12, 23, 59, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfMinute(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfMinute() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStartOfHour(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Regular Time",
			input:    time.Date(2024, 8, 12, 15, 45, 30, 0, time.UTC),
			expected: time.Date(2024, 8, 12, 15, 0, 0, 0, time.UTC),
		},
		{
			name:     "Edge Time",
			input:    time.Date(2024, 8, 12, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2024, 8, 12, 23, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfHour(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfHour() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStartOfDay(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Regular Time",
			input:    time.Date(2024, 8, 12, 15, 45, 30, 0, time.UTC),
			expected: time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Edge Time",
			input:    time.Date(2024, 8, 12, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfDay(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStartOfWeek(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Regular Time",
			input:    time.Date(2024, 8, 12, 15, 45, 30, 0, time.UTC),
			expected: time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC), // Assuming week starts on Monday
		},
		{
			name:     "Edge Time",
			input:    time.Date(2024, 8, 11, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2024, 8, 5, 0, 0, 0, 0, time.UTC), // Assuming week starts on Monday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfWeek(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfWeek() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStartOfMonth(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Regular Time",
			input:    time.Date(2024, 8, 12, 15, 45, 30, 0, time.UTC),
			expected: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Edge Time",
			input:    time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfMonth(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfMonth() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStartOfYear(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Regular Time",
			input:    time.Date(2024, 8, 12, 15, 45, 30, 0, time.UTC),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Edge Time",
			input:    time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfYear(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfYear() = %v, want %v", result, tt.expected)
			}
		})
	}
}
