package main

import (
	"fmt"
	"math/rand/v2"
)

// START OMIT

func consumer(c1, c2 chan int) int {
	var sum int
	for {
		select {
		case v, ok := <-c1:
			if !ok {
				c1 = nil
			}
			sum += v
		case v, ok := <-c2:
			if !ok {
				c2 = nil
			}
			sum += v
		default:
			if c1 == nil && c2 == nil {
				return sum
			}
		}
	}
}

// MIDDLE OMIT

func producer(c chan int) {
	for range 10 {
		c <- rand.IntN(1000)
	}

	close(c)
}

func main() {
	p1 := make(chan int)
	p2 := make(chan int)
	go producer(p1)
	go producer(p2)

	sum := consumer(p1, p2)
	fmt.Printf("Sum: %d\n", sum)
}

// END OMIT
