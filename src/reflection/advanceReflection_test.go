package reflection

import "testing"

func TestAdvanceRefection(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test AdvanceReflect"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AdvanceRefection()
		})
	}
}
