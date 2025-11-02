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
		r := changeSlice(values)
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
	for b.Loop() {
		r := changeArray(values)
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
	for b.Loop() {
		r := changeArrayReference(&values)
		if r != DefaultValue {
			b.Fatal()
		}
	}
	if values[0] != FirstValue || values[MaxValues-1] != LastValue {
		b.Fatal()
	}
}

// ARRAY END OMIT
// CHANGE START OMIT

func changeSlice(values []int) int {
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
	return values[MaxValues/2]
}

func changeArray(values [MaxValues]int) int {
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
	return values[MaxValues/2]
}

func changeArrayReference(values *[MaxValues]int) int {
	values[0] = FirstValue
	values[MaxValues-1] = LastValue
	return values[MaxValues/2]
}

// CHANGE END OMIT
