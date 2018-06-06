package routine

import (
	"testing"
)

func Benchmark_print1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print1()

	}

}

func Benchmark_goPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1()

	}
}

func Benchmark_print2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Print2()

	}

}

func Benchmark_goPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()

	}
}
func Benchmark_goPrint3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint3()

	}
}
