package web

import "testing"

func TestAdvanceWebClient(t *testing.T) {
	type args struct {
		urlname string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test www.google.co.th", args{urlname: "http://www.google.co.th"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AdvanceWebClient(tt.args.urlname)
		})
	}
}
