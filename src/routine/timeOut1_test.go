package routine

import "testing"

func TestTimeout(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestTimeout"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Timeout()
		})
	}
}
