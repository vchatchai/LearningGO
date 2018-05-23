package storage

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func TestCRUD(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	{"firstTest"},
	}
	for tt := range tests {
		fmt.Println(tt)
		// t.Run(tt.name, func(t *testing.T) {
		// 	CRUD()
		// })
	}
}
