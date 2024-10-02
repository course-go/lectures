package main

import "fmt"

func onFinish() {
	fmt.Println("Finished")
}

func main() {
	defer onFinish()

	for i := 5; i >= 0; i-- {
		fmt.Println(i)
	}

	fmt.Println("Finishing main() function")
}
