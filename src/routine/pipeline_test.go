package routine

import "testing"

func TestPipeLine(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TEST PipeLine", args{10, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PipeLine(tt.args.n1, tt.args.n2)
		})
	}
}
