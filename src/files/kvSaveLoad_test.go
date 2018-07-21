package files

import (
	"reflect"
	"testing"
)

var data string = "I Love You"

func TestSaveGOB(t *testing.T) {
	type args struct {
		file string
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test Save", args{"test.gob", data}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveGOB(tt.args.file, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SaveGOB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadGOB(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name     string
		args     args
		wantData interface{}
		wantErr  bool
	}{
		{"Test Load", args{"test.gob"}, data, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := LoadGOB(tt.args.file)
			t.Log(gotData, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGOB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("LoadGOB() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
