package main

import "fmt"

func one() (i int) {
	defer func() {
		i = 1
	}()
	return 0
}

func two() (i int) {
	defer func() {
		i += 2
	}()
	return 0
}

func main() {
	fmt.Printf("Return value of one: %d\n", one())
	fmt.Printf("Return value of two: %d\n", two())
}
