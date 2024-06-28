//go:build windows

package usignal

import (
	"context"
	"testing"
	"time"
)

func TestWithSignalCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = WithSignalCancel(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	i := 0
	for {
		select {
		case <-ctx.Done():
			if i > 3 {
				t.Errorf("i = %d, want <= 3", i)
			}
			return
		default:
			i++
			time.Sleep(time.Second)
		}
	}
}
