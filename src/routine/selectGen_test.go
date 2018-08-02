package routine

import "testing"

func TestSelectGen(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Select Gen"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectGen()
		})
	}
}
