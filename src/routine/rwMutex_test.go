package routine

import "testing"

func TestRWMutex(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test RWMutex"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RWMutex()
		})
	}
}
