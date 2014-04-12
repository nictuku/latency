package latency

import (
	"reflect"
	"testing"
	"time"
)

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

	f := make(Buckets, 0)
	for _, x := range tests {
		f.Record(x.d)
	}
	want := []int{1, 1, 2, 1, 10, 0, 0, 1, 0, 0, 1}
	if !reflect.DeepEqual([]int(f), want) {
		t.Fatalf("Bucket = %v, wanted %v", f, want)
	}
}
