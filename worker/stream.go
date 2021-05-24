package worker

import (
	"math/rand"
	"time"

	"github.com/brainly/prometheus-presentation/randx"
)

/**
Measure:
- processing time of a single event
- events batch size
- events throughput (by event type)
*/

type StreamProcessor struct{}

func NewStreamProcessor() *StreamProcessor {
	return &StreamProcessor{}
}

func (p *StreamProcessor) Work() {
	batchSize := randx.Int(1, 100)

	for i := 0; i < batchSize; i++ {
		p.processEvent(p.randomEventType())
	}
}

func (p StreamProcessor) processEvent(eventType string) {
	time.Sleep(randx.Duration(100*time.Microsecond, 500*time.Microsecond))
}

// randomEventType returns a random event type.
// There's approx. 33% chance that the type will be "A", "B" otherwise.
func (p StreamProcessor) randomEventType() string {
	if rand.Float32() < 0.33 {
		return "A"
	}

	return "B"
}
