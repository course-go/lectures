package main

import "fmt"

// START OMIT

func main() {
	c := make(chan int)
	close(c)
	value, ok := <-c
	if ok {
		fmt.Println("Value", value)
		return
	}

	fmt.Println("Channel closed")
}

// END OMIT
