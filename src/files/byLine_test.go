package files

import "testing"

func Test_lineByLine(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"TestReadLineByLine", args{"D:\\BYPASS.csv"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := lineByLine(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("lineByLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
