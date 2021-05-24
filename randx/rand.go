package randx

import (
	"math/rand"
	"time"
)

func Int(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Duration(min, max time.Duration) time.Duration {
	return time.Duration(Int(int(min), int(max)))
}
