package usafe

import (
	"context"
	"github.com/ace-zhaoy/errors"
	"log"
	"testing"
)

func TestFunc(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "TestFunc",
			args: args{
				f: func() {},
			},
			wantErr: nil,
		},
		{
			name: "TestFunc panic string",
			args: args{
				f: func() {
					panic("test error")
				},
			},
			wantErr: errors.New("test error"),
		},
		{
			name: "TestFunc panic error",
			args: args{
				f: func() {
					panic(errors.NewWithCode(10000, "test error"))
				},
			},
			wantErr: errors.NewWithCode(10000, "test error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Func(tt.args.f); !errors.Is(err, tt.wantErr) {
				t.Errorf("Func() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func Error(ctx context.Context, msg string, args ...any) {
	log.Printf(msg, args...)
}

func TestFuncWithLog(t *testing.T) {
	type args struct {
		ctx  context.Context
		logs func(ctx context.Context, msg string, args ...any)
		f    func()
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestFuncWithLog",
			args: args{
				ctx:  context.Background(),
				logs: Error,
				f:    func() { panic("test error with log") },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FuncWithLog(tt.args.ctx, tt.args.logs, tt.args.f)
		})
	}
}
