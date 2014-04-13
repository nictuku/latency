package latency

import (
	"reflect"
	"testing"
	"time"
)

// ExampleLatencyTracking creates a histogram with 10 buckets and records an event that took 16ms (16000us).
func ExampleHistogram() {
	h := &Histogram{
		Buckets:    make([]int, 10),
		Resolution: time.Millisecond,
	}
	h.Record(16 * time.Millisecond)
}

func TestLatencyTrack(t *testing.T) {
	var tests = []struct {
		d time.Duration
		b int64
	}{
		{time.Duration(1 * time.Millisecond), 0},
		{time.Duration(2 * time.Millisecond), 1},
		{time.Duration(3 * time.Millisecond), 2},
		{time.Duration(4 * time.Millisecond), 2},
		{time.Duration(8 * time.Millisecond), 3},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(16 * time.Millisecond), 4},
		{time.Duration(100 * time.Millisecond), 7},
		{time.Duration(1 * time.Second), 10},
	}

	h := &Histogram{
		Buckets:    make([]int, 0),
		Resolution: time.Millisecond,
	}
	for _, x := range tests {
		h.Record(x.d)
	}
	want := []int{1, 1, 2, 1, 10, 0, 0, 1, 0, 0, 1}
	if !reflect.DeepEqual([]int(h.Buckets), want) {
		t.Fatalf("Bucket = %v, wanted %v", h.Buckets, want)
	}
}
