package slice

import "testing"

func TestSlices(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSlice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Slices()
		})
		t.Fatal("Done.")
	}
}
