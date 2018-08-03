package contexts

import "testing"

func TestSimpleContext(t *testing.T) {
	type args struct {
		delay int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test SimpleCotext", args{delay: 4}},
		{"Test SimpleCotext", args{delay: 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SimpleContext(tt.args.delay)
		})
	}
}
