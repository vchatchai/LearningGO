package stringsutil

import "testing"

func TestRune(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Rune 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rune()
		})
	}
}
