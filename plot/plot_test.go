package plot

import (
	"time"

	"github.com/nictuku/latency"
)

// 	Example that creates a histogram with 10 buckets and records an event that took 16ms (16000us).
func ExampleLatencyTrack() {
	h := &latency.Histogram{
		Buckets:    make([]int, 0),
		Resolution: time.Millisecond,
	}
	h.Record(16 * time.Millisecond)
}
