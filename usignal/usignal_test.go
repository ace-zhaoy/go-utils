//go:build !windows
// +build !windows

package usignal

import (
	"context"
	"syscall"
	"testing"
	"time"
)

func TestWithSignalCancel(t *testing.T) {
	ctx := WithSignalCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
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
