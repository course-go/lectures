package main

import (
	"fmt"
	"time"
)

// START OMIT

func main() {
	var sum int

	go func() {
		for i := range 1_000 {
			sum += i
		}
	}()

	go func() {
		for i := range 1_000 {
			sum += i
		}
	}()

	time.Sleep(1 * time.Second) // You surely wouldn't want this in your production code.
	fmt.Printf("Summed: %d", sum)
}

// END OMIT
