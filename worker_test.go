package background

import (
	"testing"
	"time"
)

func TestWorkerStart(t *testing.T) {
	done := make(chan bool)

	// Create and start a worker
	pool := make(chan chan Task)
	w := NewWorker(1, pool)
	w.Start()
	defer w.Stop()

	// Enqueue work
	go func() {
		w.Queue <- TaskFunc(func() error {
			done <- true
			return nil
		})
	}()

	// Start processing
	<-w.Pool

	// Wait for work to be completed
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Errorf("Task Not Completed. Timeout Error")
	}
}
