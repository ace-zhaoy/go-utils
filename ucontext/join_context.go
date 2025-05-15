package glog

import (
	"context"
	"sync"
	"time"
)

type JoinContext struct {
	contexts []context.Context
	done     chan struct{}
	mu       sync.Mutex
	err      error
}

func NewJoinContext(contexts ...context.Context) (*JoinContext, context.CancelFunc) {
	jc := &JoinContext{
		contexts: contexts,
		done:     make(chan struct{}),
	}

	cancel := jc.monitor()
	return jc, cancel
}

func (jc *JoinContext) monitor() context.CancelFunc {
	var cancelOnce sync.Once
	cancel := func() {
		cancelOnce.Do(func() {
			jc.mu.Lock()
			defer jc.mu.Unlock()
			if jc.err == nil {
				jc.err = context.Canceled
				close(jc.done)
			}
		})
	}

	for _, parent := range jc.contexts {
		go func(p context.Context) {
			select {
			case <-p.Done():
				jc.mu.Lock()
				defer jc.mu.Unlock()
				if jc.err == nil {
					jc.err = p.Err()
					close(jc.done)
				}
			case <-jc.done:
				return
			}
		}(parent)
	}

	return cancel
}

func (jc *JoinContext) Deadline() (time.Time, bool) {
	var earliest time.Time
	hasDeadline := false

	for _, parent := range jc.contexts {
		if d, ok := parent.Deadline(); ok {
			if !hasDeadline || d.Before(earliest) {
				earliest = d
				hasDeadline = true
			}
		}
	}

	return earliest, hasDeadline
}

func (jc *JoinContext) Done() <-chan struct{} {
	return jc.done
}

func (jc *JoinContext) Err() error {
	jc.mu.Lock()
	defer jc.mu.Unlock()
	return jc.err
}

func (jc *JoinContext) Value(key interface{}) interface{} {
	for _, ctx := range jc.contexts {
		if value := ctx.Value(key); value != nil {
			return value
		}
	}
	return nil
}
