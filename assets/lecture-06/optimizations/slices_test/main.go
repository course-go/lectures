package optimizations_test

import "testing"

const (
	DefaultValue = 0
	FirstValue   = 69
	LastValue    = 420
	MaxValues    = 100
)

// CHANGE START OMIT

func changeMe1(values []int) { // slice
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
}

func changeMe2(values [MaxValues]int) { // array
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
}

func changeMe3(values *[MaxValues]int) { // pointer to array
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
}

// CHANGE END OMIT
// SLICE START OMIT

func BenchmarkPassSlice(b *testing.B) {
	values := make([]int, MaxValues)
	for i := 0; i < b.N; i++ {
		changeMe1(values)
	}
	if values[0] != FirstValue || values[MaxValues-1] != LastValue {
		b.Fatal()
	}
}

// SLICE END OMIT
// ARRAY START OMIT

func BenchmarkPassArrayByValue(b *testing.B) {
	values := [MaxValues]int{DefaultValue}
	for i := 0; i < b.N; i++ {
		changeMe2(values)
	}
	if values[0] != DefaultValue || values[MaxValues-1] != DefaultValue {
		b.Fatal()
	}
}

func BenchmarkPassArrayByReference(b *testing.B) {
	values := [MaxValues]int{DefaultValue}
	for i := 0; i < b.N; i++ {
		changeMe3(&values)
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
