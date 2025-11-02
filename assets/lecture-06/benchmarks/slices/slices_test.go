package slices_test

import "testing"

// START OMIT

const size = 1_000_000

func BenchmarkSliceCreation(b *testing.B) {
	for range b.N {
		slice := make([]int, 0)
		for i := range size {
			slice = append(slice, i)
		}
	}
}

func BenchmarkSliceCreationPreallocated(b *testing.B) {
	for range b.N {
		slice := make([]int, 0, size)
		for i := range size {
			slice = append(slice, i)
		}
	}
}

// END OMIT
