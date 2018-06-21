package stringsutil

import "testing"

func TestPrintStrings(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test UTF-8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintStrings()
		})
	}
}
