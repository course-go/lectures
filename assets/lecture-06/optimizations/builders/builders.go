package builders

import (
	"bytes"
	"strings"
)

// NAIVE OMIT

func NaiveBuilder(n int) string {
	var s string
	for range n {
		s += "*"
	}

	return s
}

// BYTES OMIT

func BytesBuilder(n int) string {
	var bb bytes.Buffer
	for range n {
		bb.WriteRune('*')
	}

	return bb.String()
}

// BYTES-PREALLOCATED OMIT

func BytesBuilderPreallocated(n int) string {
	var bb bytes.Buffer
	bb.Grow(n)
	for range n {
		bb.WriteRune('*')
	}

	return bb.String()
}

// STRING OMIT

func StringBuilder(n int) string {
	var sb strings.Builder
	for range n {
		sb.WriteRune('*')
	}

	return sb.String()
}

// STRING-PREALLOCATED OMIT

func StringBuilderPreallocated(n int) string {
	var sb strings.Builder
	sb.Grow(n)
	for range n {
		sb.WriteRune('*')
	}

	return sb.String()
}

// END OMIT
