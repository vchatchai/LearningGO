package files

import "testing"

func TestReadBinary(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Test ReadBinary", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadBinary(); (err != nil) != tt.wantErr {
				t.Errorf("ReadBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
