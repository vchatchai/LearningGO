package routine

import "testing"

func TestSyncGO(t *testing.T) {
	type args struct {
		totalRoutine int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test GO Routine", args{10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SyncGO(tt.args.totalRoutine)
		})
	}
}
