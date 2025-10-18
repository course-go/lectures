package main

import "fmt"

// START OMIT

func main() {
	data := make(chan int, 2)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case value := <-data:
				fmt.Println(value)
			case <-done: // Not guaranteed that all data will be read when done is signalled
				fmt.Println("Reader done!")
				return
			}
		}
	}()
	data <- 0
	data <- 1
	data <- 2
	done <- struct{}{}
	fmt.Println("All done!")
}

// END OMIT
