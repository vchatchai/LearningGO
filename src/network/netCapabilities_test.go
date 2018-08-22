package network

import "testing"

func TestNetCapabilities(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test NetCapabilities"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NetCapabilities()
		})
	}
}
