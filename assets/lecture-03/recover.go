package main

import "fmt"

func panics() {
	var i *int
	*i++
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	panics()
}
