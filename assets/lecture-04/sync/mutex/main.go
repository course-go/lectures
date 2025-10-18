package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT

func main() {
	var mu sync.Mutex
	var sum int

	go func() {
		for i := range 5 {
			mu.Lock()
			sum += i
			mu.Unlock()
		}
	}()

	go func() {
		for i := range 5 {
			mu.Lock()
			sum += i
			mu.Unlock()
		}
	}()

	time.Sleep(1 * time.Second) // You surely wouldn't want this in your production code.
	fmt.Printf("Summed: %d", sum)
}

// END OMIT
