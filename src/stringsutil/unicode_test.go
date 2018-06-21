package stringsutil

import "testing"

func TestUnitCode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Unicode"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UnitCode()
		})
	}
}
