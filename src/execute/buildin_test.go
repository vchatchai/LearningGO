package execute

import "testing"

func TestBuildinCommand(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Building Command"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BuildinCommand()
		})
	}
}
