package reflection

import "testing"

func TestReflection(t *testing.T) {
	tests := []struct {
		name string
	}{
	 {"Test Reflection"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reflection()
		})
	}
}
