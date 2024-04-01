package main

import "fmt"

// START OMIT

func compare(x, y int) bool {
	return x < y
}

func main() {
	fmt.Println(compare(1, 2))
}

// END OMIT
