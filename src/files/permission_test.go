package files

import "testing"

func TestPermission(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Permission"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Permission()
		})
	}
}
