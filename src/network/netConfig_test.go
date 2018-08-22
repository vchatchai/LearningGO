package network

import "testing"

func TestNetConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNetConfig"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NetConfig()
		})
	}
}
