package network

import "testing"

func TestDNS(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestDNS 127.0.0.1", args{input: "127.0.0.1"}},
		{"TestDNS 192.168.1.1", args{input: "192.168.1.1"}},
		{"TestDNS packtpub.com", args{input: "packtpub.com"}},
		{"TestDNS google.com", args{input: "google.com"}},
		{"TestDNS cnn.com", args{input: "cnn.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DNS(tt.args.input)
		})
	}
}
