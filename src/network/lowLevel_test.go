package network

import "testing"

func TestLowLevel(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{"LowLevel", args{"127.0.0.1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LowLevel(tt.args.address)
		})
	}
	t.Error("Done.")
}
