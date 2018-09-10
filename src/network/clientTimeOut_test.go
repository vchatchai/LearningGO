package network

import "testing"

func TestClientTimeOut(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Timeout", args{"http://www.google.co.th"}},
		{"Test Localhost", args{"http://localhost:8000"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientTimeOut(tt.args.URL)
		})
	}
}
