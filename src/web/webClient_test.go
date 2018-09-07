package web

import "testing"

func TestWebClient(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test www.google.co.th", args{url: "http://www.google.co.th"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WebClient(tt.args.url)
		})
	}
}
