package math

import (
	"fmt"
	"testing"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	fmt.Println("TestAverage")
	for _, pair := range tests {

		v := Average([]float64{1, 2})
		if v != 1.5 {
			t.Error(
				"For", pair.values,
				"expected ", pair.average,
				" got ", v)
		}
	}
}
