package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brainly/prometheus-presentation/worker"
)

func main() {
	runner := NewWorkerRunner(
		worker.NewRequestHandler(),
		worker.NewStreamProcessor(),
	)

	runner.Run()

	log.Println("all workers started")

	awaitExit()
}

func awaitExit() {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
