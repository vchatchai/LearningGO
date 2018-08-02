package routine

import "testing"

func TestTimeout2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Timeout"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Timeout2()
		})
	}
}
