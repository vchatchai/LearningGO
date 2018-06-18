package datetime

import (
	"reflect"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	type args struct {
		mytime string
	}

	tests := []struct {
		name    string
		args    args
		wantD   time.Time
		wantErr bool
	}{
		{"Test Time 17:00", args{"17:00"}, time.Now(), false},
		{"Test Time 17:00", args{"22:04"}, time.Now(), false},
	}

	ti, _ := time.Parse("15:04", "17:00")
	tests[0].wantD = ti

	ti, _ = time.Parse("15:04", "22:04")
	tests[1].wantD = ti

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := ParseTime(tt.args.mytime)
			t.Log("gotD", gotD, "err", err, " wantD:", tt.wantD)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("ParseTime() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}
