package main

import "fmt"

func sequenceID() func() int {
	var sequenceID int
	return func() int {
		sequenceID++
		return sequenceID
	}
}

func main() {
	sequenceID := sequenceID()
	fmt.Println(sequenceID())
	fmt.Println(sequenceID())
	fmt.Println(sequenceID())
}
