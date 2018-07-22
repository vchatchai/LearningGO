package osfunc

import "testing"

func TestUser(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test UserId"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			User()
		})
	}
	t.Errorf("Done.")
}
