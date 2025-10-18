package optimizations

import (
	"bytes"
	"strings"
)

// NAIVE OMIT

func naiveBuilder(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += "*"
	}

	return s
}

// BYTES OMIT

func bytesBuilder(n int) string {
	var bb bytes.Buffer
	for i := 0; i < n; i++ {
		bb.WriteRune('*')
	}

	return bb.String()
}

// BYTES-PREALLOCATED OMIT

func bytesBuilderPreallocated(n int) string {
	var bb bytes.Buffer
	bb.Grow(n)
	for i := 0; i < n; i++ {
		bb.WriteRune('*')
	}

	return bb.String()
}

// STRING OMIT

func stringBuilder(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteRune('*')
	}

	return sb.String()
}

// STRING-PREALLOCATED OMIT

func stringBuilderPreallocated(n int) string {
	var sb strings.Builder
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteRune('*')
	}

	return sb.String()
}

// END OMIT
