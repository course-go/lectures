package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Finished")
	}()

	for i := 5; i >= 0; i-- {
		fmt.Println(i)
	}

	fmt.Println("Finishing main() function")
}
