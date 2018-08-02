package routine

import "testing"

func TestBufChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Buffer Channel"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BufChannel()
		})
	}
}
