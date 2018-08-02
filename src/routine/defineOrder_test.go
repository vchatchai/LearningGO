package routine

import "testing"

func TestDefineOrder(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestDefineOrder"},
	}
	for _, tt := range tests {
		t.Log("")
		t.Run(tt.name, func(t *testing.T) {
			DefineOrder()
		})
	}
}
