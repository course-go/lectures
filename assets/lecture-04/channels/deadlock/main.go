package main

import "fmt"

func readAndPrint(c <-chan int) {
	value := <-c
	fmt.Println("Received", value)
}

func main() {
	c := make(chan int)
	c <- 5 // blocks
	go readAndPrint(c)
}
