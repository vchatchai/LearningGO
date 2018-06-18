package slice

import "testing"

func TestSortSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSortSlice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortSlice()
		})
		t.Fatal()
	}
}
