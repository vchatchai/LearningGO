package templater

import "testing"

func TestTextTemplate(t *testing.T) {
	type args struct {
		tFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test textT", args{"text.gotext"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TextTemplate(tt.args.tFile)
		})
	}
}
