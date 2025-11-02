package for_test

import "testing"

const size = 1_000_000

func seedMap() map[int]int {
	m := make(map[int]int, size)
	for i := range size {
		m[i] = i
	}

	return m
}

func seedSlice() []int {
	s := make([]int, size)
	for i := range size {
		s[i] = i
	}

	return s
}

// SLICE START OMIT

func BenchmarkSliceForRange(b *testing.B) {
	s := seedSlice()
	var sum int
	b.ResetTimer()
	for range b.N {
		for _, value := range s {
			sum += value
		}
	}
}

func BenchmarkSliceForIndex(b *testing.B) {
	s := seedSlice()
	var sum int
	b.ResetTimer()
	for range b.N {
		for i := 0; i < len(s); i++ {
			sum += s[i]
		}
	}
}

// SLICE END OMIT
// MAP START OMIT

func BenchmarkMapForLoopValue(b *testing.B) {
	m := seedMap()
	var sum int
	b.ResetTimer()
	for range b.N {
		for _, value := range m {
			sum += value
		}
	}
}

func BenchmarkMapForLoopKey(b *testing.B) {
	m := seedMap()
	var sum int
	b.ResetTimer()
	for range b.N {
		for key := range m {
			sum += m[key]
		}
	}
}

// MAP END OMIT
