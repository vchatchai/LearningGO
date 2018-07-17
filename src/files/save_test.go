package files

import "testing"

func TestSave(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Test Save()", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
