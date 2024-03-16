package main

import "fmt"

// START OMIT

func main() {
	chanA := make(chan int, 1)
	chanB := make(chan int, 1)
	chanA <- 0

	select {
	case <-chanA:
		fmt.Println("Read from A")
	case <-chanB:
		fmt.Println("Read from B")
	}
	fmt.Println("All done")
}

// END OMIT
