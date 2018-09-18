package network

import "testing"

func TestUDPClient(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUDPClient", args{"127.0.0.1:8001"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UDPClient(tt.args.address)
		})
	}
}
