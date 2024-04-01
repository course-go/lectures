package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

// START OMIT

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	data := make(chan int, 4)
	producerDone := make(chan bool)
	consumerDone := make(chan bool)

	go func() { // producer
		for {
			select {
			case data <- rand.IntN(1000):
			case <-ctx.Done():
				fmt.Println("Producer cancelled!")
				producerDone <- true
				return
			}
		}
	}()

	// MIDDLE OMIT

	go func() { // consumer
		for {
			select {
			case v := <-data:
				fmt.Println("Procesing...")
				time.Sleep(time.Second)
				fmt.Printf("Processed: %s\n", v)
			case <-ctx.Done():
				fmt.Println("Consumer cancelled!")
				consumerDone <- true
				return
			}
		}
	}()

	<-producerDone
	<-consumerDone
	fmt.Println("Program terminated...")
}

// END OMIT
