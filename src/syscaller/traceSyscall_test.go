package syscaller

import "testing"

func TestTrace(t *testing.T) {
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test", args{"ls", []string{"/tmp/"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Trace(tt.args.command, tt.args.args...)
			t.Errorf("Test")
		})
	}
}
