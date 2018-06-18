package datetime

import (
	"reflect"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	type args struct {
		myDate string
	}
	tests := []struct {
		name     string
		args     args
		wantDate time.Time
		wantErr  bool
	}{
		{"Test Date 1", args{"18 June 2018"}, time.Now(), false},
	}

	tests[0].wantDate, _ = time.Parse("02 January 2006", tests[0].args.myDate)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate, err := ParseDate(tt.args.myDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDate, tt.wantDate) {
				t.Errorf("ParseDate() = %v, want %v", gotDate, tt.wantDate)
			}
		})
	}
}
