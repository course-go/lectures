package main

import "fmt"

func main() {
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j+3 {
				fmt.Println()
				continue outer
			}

			fmt.Printf("%d%d ", i, j)
		}

		fmt.Println()
	}
}
