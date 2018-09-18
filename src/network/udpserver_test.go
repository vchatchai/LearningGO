package network

import "testing"

func TestUDPServer(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUDPServer", args{"127.0.0.1:8001"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UDPServer(tt.args.address)
		})
	}
}
