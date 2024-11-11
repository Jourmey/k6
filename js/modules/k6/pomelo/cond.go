package pomelo

import (
	"time"
)

// Placeholder is a placeholder object that can be used globally.
var Placeholder PlaceholderType

type (
	// AnyType can be used to hold any type.
	AnyType = any
	// PlaceholderType represents a placeholder type.
	PlaceholderType = struct{}
)

// A Cond is used to wait for conditions.
type Cond struct {
	signal chan PlaceholderType
}

// NewCond returns a Cond.
func NewCond() *Cond {
	return &Cond{
		signal: make(chan PlaceholderType),
	}
}

// WaitWithTimeout wait for signal return remain wait time or timed out.
func (cond *Cond) WaitWithTimeout(timeout time.Duration) (time.Duration, bool) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	begin := time.Now()
	select {
	case <-cond.signal:
		elapsed := time.Since(begin)
		remainTimeout := timeout - elapsed
		return remainTimeout, true
	case <-timer.C:
		return 0, false
	}
}

// Wait waits for signals.
func (cond *Cond) Wait() {
	<-cond.signal
}

// Signal wakes one goroutine waiting on c, if there is any.
func (cond *Cond) Signal() {
	select {
	case cond.signal <- Placeholder:
	default:
	}
}
