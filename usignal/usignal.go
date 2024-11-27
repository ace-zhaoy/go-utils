package usignal

import (
	"context"
	"log"
	"os"
	"os/signal"
)

var (
	LogEnable = true
)

func WithSignalCancel(ctx context.Context) context.Context {
	ctx, _ = WithSignalContext(ctx)
	return ctx
}

func WithSignalContext(ctx context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, Signals...)

	go func() {
		defer signal.Stop(ch)
		defer close(ch)

		select {
		case <-ctx.Done():
			if LogEnable {
				log.Printf("parent context canceled")
			}
		case s := <-ch:
			if LogEnable {
				log.Printf("got signal %s", s)
			}
			cancel()
		}
	}()
	return ctx, cancel
}
