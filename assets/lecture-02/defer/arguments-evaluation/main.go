package main

import "fmt"

// START OMIT

func function(i int) {
	fmt.Println("Deferred:", i)
}

func main() {
	x := 0
	fmt.Println("Current:", x)
	defer function(x)

	x++
	fmt.Println("Current:", x)
	defer function(x)

	x++
	fmt.Println("Current:", x)
	defer function(x)
}

// END OMIT
