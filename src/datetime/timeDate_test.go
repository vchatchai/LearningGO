package datetime

import "testing"

func TestTimeDate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test timedate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TimeDate()
		})
	}
	t.Fatal("Don.")
}
