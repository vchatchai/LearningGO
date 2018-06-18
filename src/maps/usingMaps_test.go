package maps

import "testing"

func TestUsingMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Map1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UsingMap()
		})
		t.Fatal("Done")
	}
}
