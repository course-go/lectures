package main

import (
	"fmt"
	"time"
)

// START OMIT

func process(done chan<- struct{}) { // done is write-only channel
	fmt.Println("Processing...")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished!")
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go process(done)

	fmt.Println("Waiting for processing...")
	<-done // Blocks until `process` finishes
	fmt.Println("Continuing in main")
}

// END OMIT
