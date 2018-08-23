package network

import "testing"

func TestNSrecords(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TesstNSrecode", args{domain: "mtsoukalos.eu"}},
		{"TesstNSrecode", args{domain: "www.mtsoukalos.eu"}},
		{"TesstNSrecode", args{domain: "grouplease.co.th"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NSrecords(tt.args.domain)
		})
	}
}
