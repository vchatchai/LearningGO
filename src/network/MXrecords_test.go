package network

import "testing"

func TestMXRecord(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TEST MX grouplease.co.th", args{domain: "grouplease.co.th"}},

		{"TEST MX golang.com", args{domain: "golang.com"}},
		{"TEST MX google.com", args{domain: "google.com"}},
		{"TEST MX gmail.com", args{domain: "gmail.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MXRecord(tt.args.domain)
		})
	}
}
