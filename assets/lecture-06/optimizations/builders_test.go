package optimizations_test

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkNaiveBuilder(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += "*"
	}
}

// BYTES OMIT

func BenchmarkBytesBuilder(b *testing.B) {
	var bb bytes.Buffer
	for i := 0; i < b.N; i++ {
		bb.WriteRune('*')
	}
}

// BYTES-PREALLOCATED OMIT

func BenchmarkBytesBuilderPreallocated(b *testing.B) {
	var bb bytes.Buffer
	bb.Grow(b.N)
	for i := 0; i < b.N; i++ {
		bb.WriteRune('*')
	}
}

// STRING OMIT

func BenchmarkStringBuilder(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteRune('*')
	}
}

// STRING-PREALLOCATED OMIT

func BenchmarkStringBuilderPreallocated(b *testing.B) {
	var sb strings.Builder
	sb.Grow(b.N)
	for i := 0; i < b.N; i++ {
		sb.WriteRune('*')
	}
}
