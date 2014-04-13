// Package latency provides a simple latency histogram.
package latency

import (
	"fmt"
	"math"
	"os"
	"time"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
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

// Plot saves an image of the latency histogram to filePath. The extension of filePath defines
// the format to be used - png, svg, etc.
// TODO(nictuku): Move to a sub-package to avoid import bloat.
func (h *Histogram) Plot(filePath string) error {
	count := len(h.Buckets)
	xys := make(plotter.XYs, count)

	for bucket, freq := range h.Buckets {
		xys[bucket].X = float64(bucket)
		xys[bucket].Y = float64(freq)
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("error generating plot: %v", err)
	}
	p.Title.Text = "Latency histogram"
	p.X.Label.Text = fmt.Sprintf("Latency (%v resolution)", h.Resolution)
	p.Y.Label.Text = "Frequency"

	hh, err := plotter.NewHistogram(xys, count)
	if err != nil {
		return fmt.Errorf("error generating histogram: %v", err)

	}
	p.Add(hh)

	// Save the plot to a file. Units in inches (one inch == 72 points).
	fmt.Fprintf(os.Stderr, "Saving latency histogram to %v\n", filePath)
	return p.Save(8, 6, filePath)
}
