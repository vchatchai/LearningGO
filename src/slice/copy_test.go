package slice

import "testing"

func TestCopy(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test copy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Copy()
			t.Fatal("")
		})
	}
}
