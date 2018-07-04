package reflection

import "testing"

func TestTagValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Tag"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TagValue()
		})
	}
}
