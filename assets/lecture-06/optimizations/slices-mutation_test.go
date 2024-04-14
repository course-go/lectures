package optimizations_test

// const (
// 	MAX_VALS      = 100
// 	FIRST_VALUE   = 69
// 	LAST_VALUE    = 420
// 	DEFAULT_VALUE = 0
// )

// // CHANGE START OMIT

// func changeMe1(values []int) int { // slice
// 	values[0] = FIRST_VALUE
// 	values[MAX_VALS-1] = LAST_VALUE
// 	return values[MAX_VALS/2]
// }

// func changeMe2(values [MAX_VALS]int) int { // array
// 	values[0] = FIRST_VALUE
// 	values[MAX_VALS-1] = LAST_VALUE
// 	return values[MAX_VALS/2]
// }

// func changeMe3(values *[MAX_VALS]int) int { // pointer to array
// 	values[0] = FIRST_VALUE
// 	values[MAX_VALS-1] = LAST_VALUE
// 	return values[MAX_VALS/2]
// }

// // CHANGE END OMIT
// // SLICE START OMIT

// func BenchmarkPassSlice(b *testing.B) {
// 	values := make([]int, MAX_VALS)
// 	for i := 0; i < b.N; i++ {
// 		r := changeMe1(values)
// 		if r != DEFAULT_VALUE {
// 			b.Fatal()
// 		}
// 	}
// 	if values[0] != FIRST_VALUE || values[MAX_VALS-1] != LAST_VALUE {
// 		b.Fatal()
// 	}
// }

// // SLICE END OMIT
// // ARRAY START OMIT

// func BenchmarkPassArrayByValue(b *testing.B) {
// 	values := [MAX_VALS]int{DEFAULT_VALUE}
// 	for i := 0; i < b.N; i++ {
// 		r := changeMe2(values)
// 		if r != DEFAULT_VALUE {
// 			b.Fatal()
// 		}
// 	}
// 	if values[0] != DEFAULT_VALUE || values[MAX_VALS-1] != DEFAULT_VALUE {
// 		b.Fatal()
// 	}
// }

// func BenchmarkPassArrayByReference(b *testing.B) {
// 	values := [MAX_VALS]int{DEFAULT_VALUE}
// 	for i := 0; i < b.N; i++ {
// 		r := changeMe3(&values)
// 		if r != DEFAULT_VALUE {
// 			b.Fatal()
// 		}
// 	}
// 	if values[0] != FIRST_VALUE || values[MAX_VALS-1] != LAST_VALUE {
// 		b.Fatal()
// 	}
// }

// // ARRAY END OMIT
