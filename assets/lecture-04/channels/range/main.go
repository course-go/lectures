package main

import (
	"fmt"
	"time"
)

// START OMIT

func main() {
	c := make(chan int)
	go func() {
		for i := range c { // the loop never breaks unless `close` is called
			fmt.Println(i)
		}
		fmt.Println("Channel closed")
	}()

	for num := range 5 {
		c <- num * 2
	}

	close(c)
	time.Sleep(time.Second) // wait for goroutine
}

// END OMIT
