package routine

import "testing"

func TestMutex2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMutex2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mutex2()
		})
	}
}
