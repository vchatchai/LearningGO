package datetime

import "testing"

func TestUsingTime(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test UsingTime"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UsingTime()
		})
	}
	t.Fatal("Done.")
}
