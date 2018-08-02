package routine

import "testing"

func TestNilChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Nil Channel"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NilChannel()
		})
	}
}
