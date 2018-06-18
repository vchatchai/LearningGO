package constant

import "testing"

func TestConstant(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test Constant"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Constant()
		})
		t.Fatal()
	}
}
