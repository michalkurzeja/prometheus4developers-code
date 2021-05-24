package main

import (
	"time"

	"github.com/brainly/prometheus-presentation/randx"
)

/**
Measure:
- the amount of workers
- runWorker loop duration, by worker
*/

type Worker interface {
	Work()
}

type WorkerRunner struct {
	workers []Worker
}

func NewWorkerRunner(workers ...Worker) *WorkerRunner {
	return &WorkerRunner{workers: workers}
}

func (r *WorkerRunner) Run() {
	for _, w := range r.workers {
		go r.runWorker(w)
	}
}

func (r *WorkerRunner) runWorker(w Worker) {
	for {
		time.Sleep(randx.Duration(time.Millisecond, 100*time.Millisecond))
		go w.Work()
	}
}
