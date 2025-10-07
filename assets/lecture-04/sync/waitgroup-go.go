package main

import (
	"fmt"
	"sync"
)

// START OMIT

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var sum int

	wg.Go(func() {
		for i := range 5 {
			mu.Lock()
			sum += i
			mu.Unlock()
		}
	})

	wg.Go(func() {
		for i := range 5 {
			mu.Lock()
			sum += i
			mu.Unlock()
		}
	})

	wg.Wait() // Better...
	fmt.Printf("Summed: %d", sum)
}

// END OMIT
