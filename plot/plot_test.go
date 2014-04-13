package plot

import (
	"time"

	"github.com/nictuku/latency"
)

// ExamplePlot creates a histogram with 10 buckets and records an event that took 16ms (16000us),
// then plots a (very uninteresting) graph for it.
func ExamplePlot() {
	h := &latency.Histogram{
		Buckets:    make([]int, 10),
		Resolution: time.Millisecond,
	}
	h.Record(16 * time.Millisecond)
	Plot(h, "HTTP server GET latency histogram", "test.svg")
}
