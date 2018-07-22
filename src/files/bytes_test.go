package files

import "testing"

func TestBytes(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Writer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bytes()
		})
	}
}
