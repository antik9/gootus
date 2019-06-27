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
}

func (e *Executor) awaitLaunchedTasks(launcher *Launcher, out <-chan error) {
	launcher.Lock()
	launcher.done = true
	launcher.Unlock()

	for i := 0; i < launcher.launched-e.processed; {
		if _, ok := <-out; ok {
			i++
		}
	}
}

func (l *Launcher) launchTasks(fs []func() error, in chan<- func() error) {
	for _, f := range fs {
		l.Lock()
		if l.done {
			close(in)
			l.Unlock()
			break
		} else {
			l.launched++
			l.Unlock()
			in <- f
		}
	}
}

func (l *Launcher) launchWorkers(
	workers int, processed *int, in <-chan func() error, out chan<- error) {

	workersAreDone := false
	for i := 0; i < workers; i++ {
		go func() {
			for {
				closeOnThisLoop := false
				if f, ok := <-in; ok {
					result := f()
					l.Lock()
					if l.done && l.launched-*processed == 0 && !workersAreDone {
						closeOnThisLoop = true
						workersAreDone = true
					}
					l.Unlock()
					out <- result
					if closeOnThisLoop {
						close(out)
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

func (e *Executor) processErrors(awaitedNumber int, out <-chan error) {
	for e.processed < awaitedNumber {
		if err, ok := <-out; !ok {
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
	in := make(chan func() error, 1)
	out := make(chan error, 1)
	launcher := Launcher{}

	go launcher.launchWorkers(e.maxWorkers, &e.processed, in, out)
	go launcher.launchTasks(fs, in)

	e.processErrors(len(fs), out)
	e.awaitLaunchedTasks(&launcher, out)
}
