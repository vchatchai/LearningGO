package slice

import "testing"

func TestReslicing(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "TestReslice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reslicing()
			t.Fatal("done")
		})
	}
}
