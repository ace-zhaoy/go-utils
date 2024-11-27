//go:build windows

package usignal

import (
	"os"
	"syscall"
)

var (
	Signals = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL}
)
