package worker

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/brainly/prometheus-presentation/randx"
)

/**
Measure:
- response time
- request rate (by status)
- number of concurrent requests
*/

type RequestHandler struct{}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

func (h *RequestHandler) Work() {
	time.Sleep(randx.Duration(time.Millisecond, 300*time.Millisecond))
	status := h.randomStatus()
	_ = status
}

// randomStatus returns a random status code.
// There's approx. 5% chance that the code will be "400", "200" otherwise.
func (h RequestHandler) randomStatus() int {
	if rand.Float32() < 0.05 {
		return http.StatusBadRequest
	}

	return http.StatusOK
}
