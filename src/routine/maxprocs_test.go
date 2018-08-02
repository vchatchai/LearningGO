package routine

import "testing"

func Test_getGOMAXPROCS(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test GOMAXPROCS", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getGOMAXPROCS()
			t.Logf("GOMAXPROCS %v", got)
			if got != tt.want {
				t.Errorf("getGOMAXPROCS() = %v, want %v", got, tt.want)
			}
		})
	}
}
