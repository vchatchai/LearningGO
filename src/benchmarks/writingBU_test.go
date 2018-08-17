package benchmarks

import (
	"fmt"
	"os"
	"testing"
)

func TestWritingBU(t *testing.T) {
	type args struct {
		bufferSize int
		fileSize   int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestWritingBU 1", args{bufferSize: 1, fileSize: 100000}},
		{"TestWritingBU 10", args{bufferSize: 10, fileSize: 100000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WritingBU(tt.args.bufferSize, tt.args.fileSize)
		})
	}
}

var ERR error

func benchmarkCreate(b *testing.B, buffer, filesize int) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Create("/tmp/random", buffer, filesize)
	}
	ERR = err
	err = os.Remove("/tmp/random")
	if err != nil {
		fmt.Println(err)
	}

}

func Benchmark1Create(b *testing.B) {
	benchmarkCreate(b, 1, 1000000)
}

func Benchmark2Create(b *testing.B) {
	benchmarkCreate(b, 2, 1000000)
}
func Benchmark4Create(b *testing.B) {
	benchmarkCreate(b, 4, 1000000)
}
func Benchmark10Create(b *testing.B) {
	benchmarkCreate(b, 10, 1000000)
}
func Benchmark1000Create(b *testing.B) {
	benchmarkCreate(b, 1000, 1000000)
}
