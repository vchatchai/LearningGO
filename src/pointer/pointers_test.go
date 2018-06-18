package pointer

import "testing"

func TestPointer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Pointer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pointer()
		})
	}
	t.Fatal("Done")
}
