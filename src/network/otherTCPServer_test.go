package network

import "testing"

func TestOtherTCPServer(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TEST otherTCPServer", args{":8001"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OtherTCPServer(tt.args.address)
		})
	}
}
