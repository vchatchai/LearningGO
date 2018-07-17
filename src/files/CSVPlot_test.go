package files

import "testing"

func TestCSVPlot(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test CSVPlot", args{"data"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CSVPlot(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("CSVPlot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
