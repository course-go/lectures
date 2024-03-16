package main

import "fmt"

// START OMIT

func main() {
	c := make(chan int)
	close(c)
	fmt.Println("Value", <-c)
}

// END OMIT
