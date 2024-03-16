package main

import "fmt"

func readAndPrint(c chan int) {
	value := <-c
	fmt.Println("Received", value)
}

func main() {
	c := make(chan int)
	fmt.Println("Channel length:", len(c))
	fmt.Println("Channel capacity:", cap(c))

	go readAndPrint(c)
	c <- 5
}
