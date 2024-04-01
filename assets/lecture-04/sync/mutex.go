package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// START OMIT

func main() {
	var mu sync.Mutex
	var sum int

	go func() {
		for range 5 {
			mu.Lock()
			sum += rand.IntN(100)
			mu.Unlock()
		}
	}()

	go func() {
		for range 5 {
			mu.Lock()
			sum += rand.IntN(100)
			mu.Unlock()
		}
	}()

	time.Sleep(3) // This is inappropriate, let's fix it using WaitGroup later...
	fmt.Printf("Summed: %d", sum)
}

// END OMIT
