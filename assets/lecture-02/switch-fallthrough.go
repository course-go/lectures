package main

import "fmt"

func main() {
	switch i := 5; i {
	case 4:
		fmt.Println(i, "= 4")
	case 5:
		fmt.Println(i, "= 5")
		fallthrough
	case 6:
		fmt.Println(i, "= 6")
	default:
		fmt.Println(i, "didn't match any rule")
	}
}
