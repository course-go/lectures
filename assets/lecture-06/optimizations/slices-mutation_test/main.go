package optimizations_test

import "testing"

const (
	DefaultValue = 0
	FirstValue   = 69
	LastValue    = 420
	MaxValues    = 100
)

// CHANGE START OMIT

func changeMe1(values []int) int { // slice
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
	return values[MaxValues/2]
}

func changeMe2(values [MaxValues]int) int { // array
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
	return values[MaxValues/2]
}

func changeMe3(values *[MaxValues]int) int { // pointer to array
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
	return values[MaxValues/2]
}

// CHANGE END OMIT
// SLICE START OMIT

func BenchmarkPassSlice(b *testing.B) {
	values := make([]int, MaxValues)
	for i := 0; i < b.N; i++ {
		r := changeMe1(values)
		if r != DefaultValue {
			b.Fatal()
		}
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
		r := changeMe2(values)
		if r != DefaultValue {
			b.Fatal()
		}
	}
	if values[0] != DefaultValue || values[MaxValues-1] != DefaultValue {
		b.Fatal()
	}
}

func BenchmarkPassArrayByReference(b *testing.B) {
	values := [MaxValues]int{DefaultValue}
	for i := 0; i < b.N; i++ {
		r := changeMe3(&values)
		if r != DefaultValue {
			b.Fatal()
		}
	}
	if values[0] != FirstValue || values[MaxValues-1] != LastValue {
		b.Fatal()
	}
}

// ARRAY END OMIT
