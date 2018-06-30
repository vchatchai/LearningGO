package syscaller

import "testing"

func TestSysCallUnixCmd(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test System Call Command"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SysCallUnixCmd()
		})
	}
}
