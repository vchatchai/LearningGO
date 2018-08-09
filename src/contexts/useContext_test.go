package contexts

import "testing"

func TestUseContext(t *testing.T) {
	type args struct {
		url   string
		delay int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUseContext", args{url: "http://localhost:9000", delay: 10}},
		{"TestUseContext", args{url: "http://localhost:9000", delay: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseContext(tt.args.url, tt.args.delay)
		})
	}
}
