package files

import "testing"

func TestRandom(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Teste Random", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Random(); (err != nil) != tt.wantErr {
				t.Errorf("Random() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
