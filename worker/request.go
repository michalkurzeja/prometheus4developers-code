package worker

import (
	"time"

	"github.com/brainly/prometheus-presentation/randx"
)

/**
Measure:
- response time
- request rate
- number of concurrent requests
*/

type RequestHandler struct{}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

func (h *RequestHandler) Work() {
	time.Sleep(randx.Duration(time.Millisecond, 100*time.Millisecond))
}
