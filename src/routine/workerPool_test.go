package routine

import (
	"runtime"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	type args struct {
		workers int
		jobs    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestWorkerPool", args{runtime.GOMAXPROCS(0), 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WorkerPool(tt.args.workers, tt.args.jobs)
		})
	}
}
