package main

import (
	"fmt"
	"time"
)

// START OMIT

func readAndPrint(c <-chan int) {
	value := <-c
	fmt.Println("Received", value)
}

func main() {
	c := make(chan int, 1)
	fmt.Println("Channel length:", len(c))
	fmt.Println("Channel capacity:", cap(c))

	c <- 5 // note that now it does not block
	fmt.Println("Channel length:", len(c))
	fmt.Println("Channel capacity:", cap(c))

	go readAndPrint(c)
	time.Sleep(time.Second) // need to wait
}

// END OMIT
