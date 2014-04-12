/*
Package latency provides an easy way to create latency histograms and generate graphs.

Example that creates a histogram with 10 buckets and records an event that took 16ms (16000us).

b := make(Buckets, 20)

b.Record(16 * time.Millisecond)

b.Plot("test.svg"))

20 buckets is good for tracking latency from 0us to 1s. Buckets are created in a scale of powers of
2. Even if too few buckets are specified, new buckets will be added as needed.

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

// Resolution can be changed if the latency tracking should happen at lower or right resolution than
// the default of 1 Millisecond.
// TODO(nictuku): Merge with Buckets.
var Resolution = time.Millisecond

type Buckets []int

func (b *Buckets) Record(d time.Duration) {
	var bucket int
	if d < Resolution {
		// Avoid getting to log2(0) = -Inf.
		bucket = 0
	} else {
		f := float64(d / Resolution)
		bucket = int(math.Ceil(math.Log2(f)))
		if len(*b) < bucket+1 {
			// Extend the slice.
			*b = append(*b, make(Buckets, bucket+1-len(*b))...)
		}
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
	p.X.Label.Text = fmt.Sprintf("Latency (%v resolution)", Resolution)
	p.Y.Label.Text = "Frequency"

	h, err := plotter.NewHistogram(xys, count)
	if err != nil {
		return fmt.Errorf("error generating histogram: %v", err)

	}
	p.Add(h)

	// Save the plot to a file. Units in inches (one inch == 72 points).
	fmt.Fprintf(os.Stderr, "Saving latency histogram to %v\n", filePath)
	return p.Save(8, 6, filePath)
}
