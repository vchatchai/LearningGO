package files

import (
	"testing"
)

func TestStrRead(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Read String Reader"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StrRead()
		})
	}
}

func TestStrWrite(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Read String Writer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StrWrite()
		})
	}
}
