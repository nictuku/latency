// Package latency provides a simple latency histogram.
package latency

import (
	"math"
	"time"
)

// Histogram can track latency distributions.
type Histogram struct {
	// All fields are exported to allow direct usage by the plotting library.

	// Buckets where to track the latency distribution.
	// 20 buckets are good for tracking latency from 0us to 1s.
	// Buckets are created in a scale of powers of 2. If too few buckets are
	// specified, new buckets will be added as needed.
	Buckets []int
	// Resolution of the smallest latency bucket. It must be set with
	// a positive duration.
	Resolution time.Duration
}

// Record the event in the latency histogram.
func (h *Histogram) Record(d time.Duration) {
	if d == 0 {
		panic("Histogram Resolution must be set with positive time.Duration")
	}
	var bucket int
	if d < h.Resolution {
		// Avoid getting to log2(0) = -Inf.
		bucket = 0
	} else {
		f := float64(d / h.Resolution)
		bucket = int(math.Ceil(math.Log2(f)))
		if len(h.Buckets) < bucket+1 {
			// Extend the slice.
			h.Buckets = append(h.Buckets, make([]int, bucket+1-len(h.Buckets))...)
		}
	}
	h.Buckets[bucket]++
}
