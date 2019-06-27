package parallel

import (
	"errors"
	"sync/atomic"
	"testing"
)

func TestAllTasksComplete(t *testing.T) {
	e, err := NewExecutor(5, 5)
	fs := make([]func() error, 0, 100)
	var successes, fails int64

	if err != nil {
		t.Error("Unexpected error on Executor creation")
	}

	for i := 0; i < 96; i++ {
		fs = append(fs, func() error {
			atomic.AddInt64(&successes, 1)
			return nil
		})
	}

	for i := 0; i < 4; i++ {
		fs = append(fs, func() error {
			atomic.AddInt64(&fails, 1)
			return errors.New("")
		})
	}

	e.RunTasks(fs)

	if successes != 96 {
		t.Errorf("got %d successes; want 96", successes)
	}
	if fails != 4 {
		t.Errorf("got %d fails; want 4", fails)
	}
}

func TestStopAfterNErrors(t *testing.T) {
	e, _ := NewExecutor(3, 5)
	fs := make([]func() error, 0, 100)
	var successes, fails int64

	for i := 0; i < 90; i++ {
		fs = append(fs, func() error {
			atomic.AddInt64(&fails, 1)
			return errors.New("Error")
		})
	}

	for i := 0; i < 10; i++ {
		fs = append(fs, func() error {
			atomic.AddInt64(&successes, 1)
			return nil
		})
	}

	e.RunTasks(fs)

	if fails < 5 {
		t.Errorf("got %d successes; want >= 5", fails)
	}
	if successes > 0 {
		t.Errorf("unexpectedly got some successes")
	}
}

func TestInvalidConstructor(t *testing.T) {
	_, err := NewExecutor(-1, 1)
	if err == nil {
		t.Error("invalid constuctor is allowed")
	}

	_, err = NewExecutor(1, 0)
	if err == nil {
		t.Error("invalid constuctor is allowed")
	}
}
