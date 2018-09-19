package webservice

import "testing"

func TestRestService(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestRestServiceClient"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RestService()
		})
	}
}
