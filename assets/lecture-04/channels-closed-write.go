package main

func main() {
	c := make(chan int)
	close(c)
	c <- 5
}
