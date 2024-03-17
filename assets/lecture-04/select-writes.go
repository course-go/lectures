package main

import "fmt"

// START OMIT

func main() {
	chanA := make(chan int, 1)
	chanB := make(chan int, 1)

	select {
	case chanA <- 0:
		fmt.Println("Wrote to A")
	case chanB <- 1:
		fmt.Println("Wrote to B")
	}
	fmt.Println("All done")
}

// END OMIT
