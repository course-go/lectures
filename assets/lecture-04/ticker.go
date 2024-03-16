package main

import (
	"context"
	"fmt"
	"time"
)

// START OMIT

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick!")
			case <-ctx.Done():
				fmt.Println("Ticker done!")
				done <- true
				return
			}
		}
	}()

	time.Sleep(10 * time.Second)
	cancel() // shutdown worker
	<-done   // wait for worker to clean-up
	fmt.Println("All done!")
}

// END OMIT
