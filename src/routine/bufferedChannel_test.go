package routine

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	fmt.Println("TestBufferedChannel", "Start")
	tests := []struct {
		name string
	}{
		{"Test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BufferedChannel()
		})
	}
	fmt.Println("TestBufferedChannel", "Stop")
}
