package main

import "fmt"

// START OMIT

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Println("Classified", x, "to", classify(x))
	}
}

func classify(x int) string {
	defer func(int) {
		fmt.Println("Deferred", x)
	}(x)
	switch x {
	case 0, 2, 4, 6, 8:
		return "even number"
	case 1, 3, 5, 7, 9:
		return "odd number"
	default:
		return "?"
	}
}

// END OMIT
