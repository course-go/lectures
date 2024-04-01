package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
)

// START OMIT

func main() {
	var wg sync.WaitGroup
	var sum atomic.Int64
	wg.Add(2)

	go func() {
		for range 5 {
			value := int64(rand.IntN(100))
			sum.Add(value)
		}
		wg.Done() // Decrements the WaitGroup counter
	}()
	go func() {
		for range 5 {
			value := int64(rand.IntN(100))
			sum.Add(value)
		}
		wg.Done() // Decrements the WaitGroup counter
	}()

	wg.Wait() // Better...
	fmt.Printf("Summed: %d", sum.Load())
}

// END OMIT
