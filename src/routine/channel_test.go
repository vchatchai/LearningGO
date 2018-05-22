package routine

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	fmt.Println("TestChannel")
	tests := []struct {
		name string
	}{
		{"Test 1"},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			Channel()
		})
	}
}
