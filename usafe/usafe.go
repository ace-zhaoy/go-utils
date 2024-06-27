package usafe

import (
	"context"
	"github.com/ace-zhaoy/errors"
)

func Func(f func()) (err error) {
	defer errors.Recover(func(e error) { err = errors.WithStack(e) })
	f()
	return
}

func FuncWithLog(ctx context.Context, logErrFunc func(ctx context.Context, msg string, args ...any), f func()) {
	defer errors.Recover(func(e error) { logErrFunc(ctx, "%+v", errors.WithStack(e)) })
	f()
}
