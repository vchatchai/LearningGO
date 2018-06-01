package routine

import (
	"testing"
	"time"
)

func Test_print1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"non concurrent test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			print1()
		})
	}
}

func Test_goPrint1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"concurrent test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goPrint1()
			time.Sleep(1 * time.Millisecond)
		})
	}
}
