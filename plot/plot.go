package plot

import (
	"fmt"
	"os"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"github.com/nictuku/latency"
)

// Plot saves an image of the latency histogram to filePath. The extension of filePath defines
// the format to be used - png, svg, etc.
func Plot(h *latency.Histogram, description, filePath string) error {
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
	p.Title.Text = description
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
