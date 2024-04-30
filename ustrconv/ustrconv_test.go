package ustrconv

import "testing"

func TestParseInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Normal number", args: args{s: "123"}, want: 123, wantErr: false},
		{name: "Negative number", args: args{s: "-123"}, want: -123, wantErr: false},
		{name: "Number with leading zeros", args: args{s: "00123"}, want: 123, wantErr: false},
		{name: "Empty string", args: args{s: ""}, want: 0, wantErr: true},
		{name: "Non-numeric string", args: args{s: "abc"}, want: 0, wantErr: true},
		{name: "Number beyond int range", args: args{s: "9223372036854775808"}, want: 9223372036854775807, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUint64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{name: "Normal number", args: args{s: "123"}, want: 123, wantErr: false},
		{name: "Number with leading zeros", args: args{s: "00123"}, want: 123, wantErr: false},
		{name: "Empty string", args: args{s: ""}, want: 0, wantErr: true},
		{name: "Non-numeric string", args: args{s: "abc"}, want: 0, wantErr: true},
		{name: "Negative number", args: args{s: "-123"}, want: 0, wantErr: true},
		{name: "Number beyond uint64 range", args: args{s: "18446744073709551616"}, want: 18446744073709551615, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUint64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseHexInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "Normal number", args: args{s: "123"}, want: 291, wantErr: false},
		{name: "Negative number", args: args{s: "-123"}, want: -291, wantErr: false},
		{name: "Number with leading zeros", args: args{s: "00123"}, want: 291, wantErr: false},
		{name: "Empty string", args: args{s: ""}, want: 0, wantErr: true},
		{name: "Non-numeric string", args: args{s: "abc"}, want: 2748, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseHexInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHexInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseHexInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
