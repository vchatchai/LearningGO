package benchmarks

import "testing"

func TestCannotReach(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Tesst Can't Reach"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CannotReach()
		})
	}
}
