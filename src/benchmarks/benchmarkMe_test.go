package benchmarks

import (
	"testing"
)

func Test_fibo1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TEST  FIBO1", args{40}, 102334155},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibo1(tt.args.n)
			t.Log(got)
			if got != tt.want {
				t.Errorf("fibo1() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_fibo2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TEST  FIBO2", args{40}, 102334155},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibo2(tt.args.n)
			t.Log(got)
			if got != tt.want {
				t.Errorf("fibo2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibo3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TEST  FIBO3", args{40}, 102334155},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibo3(tt.args.n)
			t.Log(got)
			if got != tt.want {
				t.Errorf("fibo3() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result int

func benchmarkfibo1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo1(n)
	}
	result = r
}

func benchmarkfibo2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo2(n)
	}
	result = r
}
func benchmarkfibo3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo3(n)
	}
	result = r
}

func Benchmark30fibo1(b *testing.B) {
	benchmarkfibo1(b, 30)
}
func Benchmark30fibo2(b *testing.B) {
	benchmarkfibo2(b, 30)
}
func Benchmark30fibo3(b *testing.B) {
	benchmarkfibo3(b, 30)
}

func Benchmark50fibo1(b *testing.B) {
	benchmarkfibo1(b, 50)
}
func Benchmark50fibo2(b *testing.B) {
	benchmarkfibo2(b, 50)
}
func Benchmark50fibo3(b *testing.B) {
	benchmarkfibo3(b, 50)
}
