package background

import (
	"testing"
	"time"
)

func TestDispatcher(t *testing.T) {
	done := make(chan bool)

	pool := make(chan Task, 100)
	dispatcher := NewDispatcher(1)
	dispatcher.Start()
	dispatcher.Dispatch(pool)

	go func() {
		pool <- TaskFunc(func() error {
			done <- true
			return nil
		})
	}()

	// Wait for work to be completed
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Errorf("Task not completed. Timeout Error")
	}
}
