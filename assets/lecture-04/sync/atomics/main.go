package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// START OMIT

func main() {
	var wg sync.WaitGroup
	var sum atomic.Int64

	wg.Go(func() {
		for i := range 5 {
			sum.Add(int64(i))
		}
	})

	wg.Go(func() {
		for i := range 5 {
			sum.Add(int64(i))
		}
	})

	wg.Wait()
	fmt.Printf("Summed: %d", sum.Load())
}

// END OMIT
