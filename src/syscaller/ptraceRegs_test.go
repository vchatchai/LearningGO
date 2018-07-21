package syscaller

import "testing"

func TestPtrace(t *testing.T) {
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"TestPtrace", args{"echo", []string{"Mastering GO!"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Ptrace(tt.args.command, tt.args.args...)
			t.Log(err)

			if (err != nil) != tt.wantErr {
				t.Errorf("Ptrace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
