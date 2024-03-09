package main

import "fmt"

func main() {
	helloWorld := "plɹoʍ ollǝɥ"

	for _, ch := range helloWorld { // iterates over runes - i.e. unicode glyphs
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	for i := 0; i < len(helloWorld); i++ { // iterates over bytes
		fmt.Printf("%v ", helloWorld[i])
	}

	fmt.Println()

	fmt.Println("Rune count:", len([]rune(helloWorld)))
	fmt.Println("Byte count:", len(helloWorld))
}
