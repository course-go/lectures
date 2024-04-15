package optimizations_test

import "testing"

const (
	MAX_VALS      = 100
	FIRST_VALUE   = 69
	LAST_VALUE    = 420
	DEFAULT_VALUE = 0
)

// CHANGE START OMIT

func changeMe1(values []int) { // slice
	values[0] = FIRST_VALUE
	values[MAX_VALS-1] = LAST_VALUE
}

func changeMe2(values [MAX_VALS]int) { // array
	values[0] = FIRST_VALUE
	values[MAX_VALS-1] = LAST_VALUE
}

func changeMe3(values *[MAX_VALS]int) { // pointer to array
	values[0] = FIRST_VALUE
	values[MAX_VALS-1] = LAST_VALUE
}

// CHANGE END OMIT
// SLICE START OMIT

func BenchmarkPassSlice(b *testing.B) {
	values := make([]int, MAX_VALS)
	for i := 0; i < b.N; i++ {
		changeMe1(values)
	}
	if values[0] != FIRST_VALUE || values[MAX_VALS-1] != LAST_VALUE {
		b.Fatal()
	}
}

// SLICE END OMIT
// ARRAY START OMIT

func BenchmarkPassArrayByValue(b *testing.B) {
	values := [MAX_VALS]int{DEFAULT_VALUE}
	for i := 0; i < b.N; i++ {
		changeMe2(values)
	}
	if values[0] != DEFAULT_VALUE || values[MAX_VALS-1] != DEFAULT_VALUE {
		b.Fatal()
	}
}

func BenchmarkPassArrayByReference(b *testing.B) {
	values := [MAX_VALS]int{DEFAULT_VALUE}
	for i := 0; i < b.N; i++ {
		changeMe3(&values)
	}
	if values[0] != FIRST_VALUE || values[MAX_VALS-1] != LAST_VALUE {
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
