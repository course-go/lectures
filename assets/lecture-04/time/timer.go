package main

import (
	"fmt"
	"time"
)

// START OMIT

func main() {
	done := make(chan bool)
	go func() {
		timer := time.NewTimer(3 * time.Second)

		fmt.Println("Waiting...")
		<-timer.C

		fmt.Println("Processing...")
		time.Sleep(2 * time.Second)

		done <- true
	}()

	<-done // wait for worker
	fmt.Println("All done!")
}

// END OMIT
