package benchmarks

import (
	"testing"
)

var r int

func BenchmarkForloop30(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = forloop(30)
	}
	r = result
}

func BenchmarkForloop50(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = forloop(50)
	}
	r = result
}

func BenchmarkRecursive30(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		recursive(30)
	}
	r = result
}

func BenchmarkRecursive50(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		recursive(50)
	}
	r = result
}
