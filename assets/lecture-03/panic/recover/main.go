package main

import "fmt"

// START OMIT

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	increment(nil)
}

func increment(i *int) {
	*i++
}

// END OMIT
