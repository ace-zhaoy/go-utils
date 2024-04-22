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
