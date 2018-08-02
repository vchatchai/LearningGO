package routine

import "testing"

func TestChannelOfChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"ChannelOfChannel"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChannelOfChannel()
		})
	}
}
