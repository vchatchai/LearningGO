package network

import "testing"

func TestAnotherTimeOut(t *testing.T) {
	type args struct {
		URL    string
		second int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUnitTest", args{"http://localhost:8001/time", 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AnotherTimeOut(tt.args.URL, tt.args.second)
		})
	}
}
