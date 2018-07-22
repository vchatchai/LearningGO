package routine

import "testing"

func TestWriteCH(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Channel"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteCH()
		})
	}
}
