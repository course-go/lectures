package main

func main() {
	var channel chan int = make(chan int)
	channel <- 5 // note that this deadlocks
	_ = <-channel
}
