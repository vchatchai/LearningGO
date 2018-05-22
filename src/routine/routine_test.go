package routine

import (
	"fmt"
	"testing"
)

func Test_F2(t *testing.T) {
	fmt.Println("Test_f2")
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Case 1", args{10}},
		{"Case 2", args{10}},
	}
	for k, tt := range tests {
		fmt.Println(k, tt.name)
		t.Run(tt.name, func(t *testing.T) {
			F2(tt.args.n)
		})
	}
}
