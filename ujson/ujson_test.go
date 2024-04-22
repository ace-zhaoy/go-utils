package ujson

import "testing"

func TestToJson(t *testing.T) {
	type args struct {
		V string `json:"v"`
		I int    `json:"i"`
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{V: "test", I: 1},
			want: `{"v":"test","i":1}`,
		},
		{
			name: "你好",
			args: args{V: "你好", I: 2},
			want: `{"v":"你好","i":2}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToJson(tt.args); got != tt.want {
				t.Errorf("ToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
