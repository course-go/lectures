package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

// START OMIT

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var sum int
	wg.Add(2)
	go func() {
		for range 5 {
			mu.Lock()
			sum += rand.IntN(100)
			mu.Unlock()
		}
		wg.Done() // Decrements the WaitGroup counter
	}()
	go func() {
		for range 5 {
			mu.Lock()
			sum += rand.IntN(100)
			mu.Unlock()
		}
		wg.Done() // Decrements the WaitGroup counter
	}()
	wg.Wait() // Better...
	fmt.Printf("Summed: %d", sum)
}

// END OMIT
