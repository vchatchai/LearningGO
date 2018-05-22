package routine

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	fmt.Println("TestSelect start")
	tests := []struct {
		name string
	}{
		{"Test 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Select()
			time.Sleep(time.Second * 10)
		})
	}

	fmt.Println("TestSelect done.")
}
