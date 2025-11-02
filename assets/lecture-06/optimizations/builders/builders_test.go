package builders_test

import (
	"testing"

	"github.com/course-go/lectures/assets/lecture-06/optimizations/builders"
)

// START OMIT

func BenchmarkNaiveBuilder(b *testing.B) {
	builders.NaiveBuilder(b.N)
}

func BenchmarkBytesBuilder(b *testing.B) {
	builders.BytesBuilder(b.N)
}

func BenchmarkBytesBuilderPreallocated(b *testing.B) {
	builders.BytesBuilderPreallocated(b.N)
}

func BenchmarkStringBuilder(b *testing.B) {
	builders.StringBuilder(b.N)
}

func BenchmarkStringBuilderPreallocated(b *testing.B) {
	builders.StringBuilderPreallocated(b.N)
}

// END OMIT
