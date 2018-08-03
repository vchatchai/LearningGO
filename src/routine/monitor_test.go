package routine

import "testing"

func TestShareMemory(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test ShareMemory", args{n: 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShareMemory(tt.args.n)
		})
	}
}
