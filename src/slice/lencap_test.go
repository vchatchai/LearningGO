package slice

import "testing"

func TestLencap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "TestLenCap"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Lencap()
			t.Fatal("done.")
		})
	}
}
