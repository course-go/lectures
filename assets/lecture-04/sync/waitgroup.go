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
	wg.Add(2)
	go func() {
		for i := range 5 {
			mu.Lock()
			sum += i
			mu.Unlock()
		}
		wg.Done() // Decrements the WaitGroup counter
	}()
	go func() {
		for i := range 5 {
			mu.Lock()
			sum += i
			mu.Unlock()
		}
		wg.Done() // Decrements the WaitGroup counter
	}()
	wg.Wait()
	fmt.Printf("Summed: %d", sum)
}

// END OMIT
