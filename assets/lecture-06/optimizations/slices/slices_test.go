package slices_test

import "testing"

const (
	DefaultValue = 0
	FirstValue   = 69
	LastValue    = 420
	MaxValues    = 100
)

// SLICE START OMIT

func BenchmarkPassSlice(b *testing.B) {
	values := make([]int, MaxValues)
	for b.Loop() {
		changeSlice(values)
	}
	if values[0] != FirstValue || values[MaxValues-1] != LastValue {
		b.Fatal()
	}
}

// SLICE END OMIT
// ARRAY START OMIT

func BenchmarkPassArrayByValue(b *testing.B) {
	values := [MaxValues]int{DefaultValue}
	for b.Loop() {
		changeArray(values)
	}
	if values[0] != DefaultValue || values[MaxValues-1] != DefaultValue {
		b.Fatal()
	}
}

func BenchmarkPassArrayByReference(b *testing.B) {
	values := [MaxValues]int{DefaultValue}
	for b.Loop() {
		changeArrayReference(&values)
	}
	if values[0] != FirstValue || values[MaxValues-1] != LastValue {
		b.Fatal()
	}
}

// ARRAY END OMIT
// SLICES-PREALLOCATION START OMIT

func BenchmarkSliceInsertion(b *testing.B) {
	s := make([]int, 0)
	for i := range b.N {
		s = append(s, i)
	}
}

func BenchmarkPreallocatedSliceInsertion(b *testing.B) {
	s := make([]int, 0, b.N)
	for i := range b.N {
		s = append(s, i)
	}
}

// SLICES-PREALLOCATION END OMIT
// CHANGE START OMIT

func changeSlice(values []int) {
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
}

func changeArray(values [MaxValues]int) {
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
}

func changeArrayReference(values *[MaxValues]int) {
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
}

// CHANGE END OMIT
