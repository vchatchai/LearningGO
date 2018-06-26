package stringsutil

import "testing"

func TestUseStrings(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test UseString"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseStrings()
		})
	}
}
