package web

import "testing"

func TestHttpTrace(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test localhost:8001", args{"http://localhost:8001"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpTrace(tt.args.url)
		})
	}
}
