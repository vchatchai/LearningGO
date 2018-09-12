package network

import "testing"

func TestTCPclient(t *testing.T) {
	type args struct {
		address string
		text    string
	}
	tests := []struct {
		name string
		args args
	}{
		// {"Test TCP client", args{"127.0.0.1:8001", "hello world\n"}},
		{"Test TCP client", args{"127.0.0.1:8001", "STOP"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TCPclient(tt.args.address, tt.args.text)
		})
	}
}
