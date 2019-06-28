package parallel

import (
	"errors"
	"sync"
)

type Executor struct {
	errorCounter int
	maxErrors    int
	maxWorkers   int
	processed    int
}

type Launcher struct {
	sync.Mutex
	launched int
	done     bool
	tasks    chan func() error
	result   chan error
}

func (l *Launcher) launchTasks(fs []func() error) {
	for _, f := range fs {
		l.Lock()
		if l.done {
			l.Unlock()
			close(l.tasks)
			break
		} else {
			l.launched++
			l.Unlock()
			l.tasks <- f
		}
	}
}

func (l *Launcher) launchWorkers(workers int, processed *int) {
	workersAreDone := false
	for i := 0; i < workers; i++ {
		go func() {
			for {
				closeOnThisLoop := false
				if f, ok := <-l.tasks; ok {
					result := f()
					l.Lock()
					if l.done && l.launched-*processed == 0 && !workersAreDone {
						closeOnThisLoop = true
						workersAreDone = true
					}
					l.Unlock()
					l.result <- result
					if closeOnThisLoop {
						close(l.result)
					}
				} else {
					break
				}
			}
		}()
	}
}

func NewExecutor(maxWorkers, maxErrors int) (*Executor, error) {
	if maxErrors <= 0 || maxWorkers <= 0 {
		return nil, errors.New("provide positive parameters")
	}
	return &Executor{
		maxErrors:  maxErrors,
		maxWorkers: maxWorkers}, nil
}

func NewLauncher() *Launcher {
	return &Launcher{
		tasks:  make(chan func() error),
		result: make(chan error)}
}

func (e *Executor) processErrors(awaitedNumber int, launcher *Launcher) {
	for e.processed < awaitedNumber {
		if err, ok := <-launcher.result; !ok {
			continue
		} else if err != nil {
			e.errorCounter++
		}

		e.processed++
		if e.errorCounter >= e.maxErrors {
			break
		}
	}
}

func (e *Executor) RunTasks(fs []func() error) {
	launcher := NewLauncher()

	go launcher.launchWorkers(e.maxWorkers, &e.processed)
	go launcher.launchTasks(fs)

	e.processErrors(len(fs), launcher)
	e.stopLauncher(launcher)
}

func (l *Launcher) stop() {
	l.Lock()
	l.done = true
	l.Unlock()
}

func (e *Executor) stopLauncher(launcher *Launcher) {
	launcher.stop()

	for i := 0; i < launcher.launched-e.processed; {
		if _, ok := <-launcher.result; ok {
			i++
		}
	}
}
