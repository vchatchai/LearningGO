package syscaller

import "testing"

func TestSysCall(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Syscall"},
	}
	for _, tt := range tests {
		t.Log(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			SysCallUnix()
		})
	}
}
