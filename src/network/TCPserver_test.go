package network

import "testing"

func TestTCPserver(t *testing.T) {
	type args struct {
		port string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TEST server", args{":8001"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TCPserver(tt.args.port)
		})
	}
}
