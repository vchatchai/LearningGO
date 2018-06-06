package inputoutput

import "testing"

func TestStdout(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Stdout()
			t.Error("done")
		})
	}
}
