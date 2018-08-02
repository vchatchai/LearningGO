package stringsutil

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testData = []string{"a", "b", "c", "d", "e"}
)

func BenchmarkBuild(b *testing.B) {
	var build strings.Builder
	for i := 0; i < b.N; i++ {
		for _, s := range testData {
			build.WriteString(s)
			build.WriteString(":")

		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(testData, ":")
		_ = s
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s:%s:%s:%s:%s", testData[0], testData[1], testData[2], testData[3], testData[4])
		_ = s
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := testData[0] + ":" + testData[1] + ":" + testData[2] + ":" + testData[3] + ":" + testData[4]
		_ = s
	}
}
