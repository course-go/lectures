package main

import "fmt"

// START OMIT

func compareInts(x, y int) bool {
	return x < y
}

func compareFloats(x, y float64) bool {
	return x < y
}

func compareStrings(x, y string) bool {
	return x < y
}

func main() {
	fmt.Println(compareInts(1, 2))
	fmt.Println(compareFloats(1.5, 2.6))
	fmt.Println(compareStrings("Go", "Gophers!"))
}

// END OMIT
