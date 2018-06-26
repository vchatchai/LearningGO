package maps

import "testing"

func TestKeyValue(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TEST UNKNOWN", args{[]string{"UNKNOWN"}}},
		{"TEST ADD 123 1 2 3", args{[]string{"ADD", "123", "1", "2", "3"}}},
		{"TEST ADD 234 2 3 4", args{[]string{"ADD", "234", "2", "3", "4"}}},
		{"TEST ADD 345", args{[]string{"ADD", "345"}}},
		{"TEST PRINT", args{[]string{"PRINT"}}},
		{"TEST ADD 345 3 4 5", args{[]string{"ADD", "345", "3", "4", "5"}}},
		{"TEST PRINT", args{[]string{"PRINT"}}},
		{"TEST CHANGE 345 3 4 5", args{[]string{"CHANGE", "345", "3", "4", "5"}}},
		{"TEST PRINT", args{[]string{"PRINT"}}},
		{"TEST DELETE 345", args{[]string{"DELETE", "345"}}},
		{"TEST PRINT", args{[]string{"PRINT"}}},
		{"TEST ADD 345 3 4 5", args{[]string{"ADD", "345", "3", "4", "5"}}},
		{"TEST ADD 567 -5 -6 -7", args{[]string{"ADD", "567", "-5", "-6", "-7"}}},
		{"TEST CHANGE 345", args{[]string{"CHANGE", "345"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			KeyValue(tt.args.tokens...)
		})
	}
}
