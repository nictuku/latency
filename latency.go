/*
Package latency provides an easy way to create latency histograms and generate graphs.

Example that creates a histogram with 10 buckets and records an event that took 16ms.

b := make(Buckets, 10)

b.Record(16 * time.Millisecond)

b.Plot("test.svg"))

10 buckets is good for tracking latency from 0ms to 1s.
Even if too few buckets are specified, new buckets will be added as needed.

Minimum resolution: 1 millisecond.
Scale: powers of 2.
*/
package latency

import (
	"fmt"
	"math"
	"os"
	"time"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
)

type Buckets []int

func (b *Buckets) Record(d time.Duration) {
	ms := float64(d.Nanoseconds() / 1e6)
	bucket := int(math.Ceil(math.Log2(ms)))
	if len(*b) < bucket+1 {
		// Extend the slice.
		*b = append(*b, make(Buckets, bucket+1-len(*b))...)
	}
	(*b)[bucket]++
}

// Plot saves an image of the latency histogram to filePath. The extension of filePath defines
// the format to be used - png, svg, etc.
func (b *Buckets) Plot(filePath string) error {
	count := len(*b)

	xys := make(plotter.XYs, count)

	for bucket, freq := range *b {
		xys[bucket].X = float64(bucket)
		xys[bucket].Y = float64(freq)
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("error generating plot: %v", err)
	}
	p.Title.Text = "Latency histogram"
	p.X.Label.Text = "Latency (ms)"
	p.Y.Label.Text = "Frequency"

	h, err := plotter.NewHistogram(xys, count)
	if err != nil {
		return fmt.Errorf("error generating histogram: %v", err)

	}
	p.Add(h)

	// Save the plot to a file. Units in inches (one inch == 72 points).
	fmt.Fprintf(os.Stderr, "Saving latency histogram to", filePath)
	return p.Save(8, 6, filePath)
}
