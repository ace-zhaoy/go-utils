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
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal)
	signal.Notify(ch, Signals...)
	go func() {
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
	return ctx
}
