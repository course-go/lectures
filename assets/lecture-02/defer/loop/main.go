package main

import "fmt"

func onFinish(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i <= 5; i++ {
		defer onFinish(i)
	}

	fmt.Println("Finishing main() function")
}
