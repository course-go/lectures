package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ { // classic C for
		fmt.Println(i)
	}

	var i int = 10
	for i < 11 { // while with condition
		fmt.Println(i)
		i++
	}

	for { // while true
		fmt.Println("Break the loop!")
		break
	}
}
